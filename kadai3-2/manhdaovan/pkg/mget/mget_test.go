package mget_test

import (
	"context"
	"crypto/md5"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo5/kadai3-2/manhdaovan/pkg/mget"
)

// initMockServerFile inits a httptest.Server and a callback to close server
func initMockServerFile(t *testing.T, initFnc func(t *testing.T) http.HandlerFunc) (*httptest.Server, func()) {
	t.Helper()
	ts := httptest.NewServer(initFnc(t))
	return ts, func() { ts.Close() }
}

type mockServerFile struct {
	file string
}

func (msf *mockServerFile) mockServerHandler(t *testing.T) http.HandlerFunc {
	t.Helper()
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Accept-Ranges", "bytes")
		headerRange := r.Header.Get("Range")

		body := func() []byte {
			testDataBytes, err := ioutil.ReadFile(msf.file)
			if err != nil {
				t.Errorf("cannot read test file:%s, error: %v", msf.file, err)
			}

			if headerRange == "" {
				return testDataBytes
			}

			rangeItems := strings.Split(headerRange, "=")
			if rangeItems[0] != "bytes" {
				t.Errorf("range header should be bytes, got: %s", headerRange)
			}
			rangeValues := strings.Split(rangeItems[1], "-")

			rangeFrom, err := strconv.Atoi(rangeValues[0])
			if err != nil {
				t.Errorf("invalid range-from value: %v, error: %v", rangeValues[0], err)
			}

			rangeTo, err := strconv.Atoi(rangeValues[1])
			if err != nil {
				t.Errorf("invalid range-to value: %v, error: %v", rangeValues[1], err)
			}

			return testDataBytes[rangeFrom:rangeTo]
		}()

		w.Header().Set("Content-Length", strconv.Itoa(len(body)))
		w.WriteHeader(http.StatusPartialContent)
		w.Write(body)
	}
}

func TestMGet_Download(t *testing.T) {
	type fields struct {
		workerNum uint
	}
	type args struct {
		dst string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "file bytes are not divisible by workers",
			fields:  fields{workerNum: 5},
			args:    args{dst: "./testdata/out/"},
			want:    "testdata/out/378bytes.text",
			wantErr: false,
		},
		{
			name:    "file bytes are divisible by workers",
			fields:  fields{workerNum: 3},
			args:    args{dst: "./testdata/out/"},
			want:    "testdata/out/378bytes.text",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServerFile := mockServerFile{file: "./testdata/378bytes.text"}
			mockServer, shutdownMockServer := initMockServerFile(t, mockServerFile.mockServerHandler)
			defer shutdownMockServer()

			m := mget.NewMGet(
				http.DefaultClient,
				tt.fields.workerNum,
				mget.DefaultExitSigs,
				"", "",
			)
			got, err := m.Download(context.Background(), tt.args.dst, mockServer.URL+"/file/378bytes.text")
			if (err != nil) != tt.wantErr {
				t.Errorf("MGet.Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want || !isSameFileContent(t, got, tt.want) {
				t.Errorf("MGet.Download() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isSameFileContent(t *testing.T, file1, file2 string) bool {
	t.Helper()
	bytesFile1, err := ioutil.ReadFile(file1)
	if err != nil {
		t.Errorf("cannot open file: %s", file1)
		return false
	}
	bytesFile2, err := ioutil.ReadFile(file2)
	if err != nil {
		t.Errorf("cannot open file: %s", file2)
		return false
	}

	return md5.Sum(bytesFile1) == md5.Sum(bytesFile2)
}
