package mypget_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo5/kadai3-2/lfcd85/mypget"
)

type testServerFile struct {
	testDataPath string
}

func TestDownloader_Execute(t *testing.T) {
	cases := []struct {
		testDataPath string
		splitNum     int
	}{
		{"../testdata/tower.jpg", 8},
		{"../testdata/lorem_ipsum.txt", 4},
	}

	for _, c := range cases {
		c := c
		t.Run(c.testDataPath, func(t *testing.T) {
			tsf := testServerFile{c.testDataPath}

			ts, closeTs := initTestServer(t, tsf.testServerHandler)
			defer closeTs()

			d := initDownloader(t, ts.URL, c.splitNum)
			if err := d.Execute(nil); err != nil {
				t.Errorf("failed to execute split downloader: %v", err)
			}
			deleteOutputFile(t, d)
		})
	}
}

func initDownloader(t *testing.T, urlStr string, splitNum int) *mypget.Downloader {
	t.Helper()

	url, err := url.Parse(urlStr)
	if err != nil {
		t.Errorf("failed to parse URL for testing: %v", err)
	}

	return mypget.New(url, splitNum)
}

func deleteOutputFile(t *testing.T, d *mypget.Downloader) {
	t.Helper()

	outputPath := d.ExportOutputPath()
	if outputPath != "" {
		if err := os.Remove(outputPath); err != nil {
			t.Errorf("failed to remove output file: %v", err)
		}
	}
}

func initTestServer(t *testing.T, handler func(t *testing.T, w http.ResponseWriter, r *http.Request)) (*httptest.Server, func()) {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			handler(t, w, r)
		},
	))

	return ts, func() { ts.Close() }
}

func (tsf *testServerFile) testServerHandler(t *testing.T, w http.ResponseWriter, r *http.Request) {
	t.Helper()

	w.Header().Set("Accept-Ranges", "bytes")
	headerRange := r.Header.Get("Range")

	body := func() []byte {
		testDataBytes, err := ioutil.ReadFile(tsf.testDataPath)
		if err != nil {
			t.Errorf("failed to read the test file in test server: %v", err)
		}

		if headerRange == "" {
			return testDataBytes
		}

		rangeItems := strings.Split(headerRange, "=")
		if rangeItems[0] != "bytes" {
			t.Errorf("range in test server should have bytes value, but actually does not")
		}
		rangeValues := strings.Split(rangeItems[1], "-")

		rangeStart, err := strconv.Atoi(rangeValues[0])
		if err != nil {
			t.Errorf("failed to get the start of the range in test server: %v", err)
		}

		rangeEnd, err := strconv.Atoi(rangeValues[1])
		if err != nil {
			t.Errorf("failed to get the end of the range in test server: %v", err)
		}

		return testDataBytes[rangeStart:rangeEnd]
	}()

	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(http.StatusPartialContent)

	if _, err := w.Write(body); err != nil {
		t.Errorf("failed to write the response body in test server: %v", err)
	}
}
