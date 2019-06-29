package omikuji

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var cases = []struct {
	name string
	time time.Time
}{
	{
		name: "case1",
		time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		name: "case2",
		time: time.Date(2019, 1, 1, 23, 59, 59, 0, time.UTC),
	},
	{
		name: "case3",
		time: time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC),
	},
	{
		name: "case4",
		time: time.Date(2019, 1, 2, 23, 59, 59, 0, time.UTC),
	},
	{
		name: "case5",
		time: time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
	},
	{
		name: "case6",
		time: time.Date(2019, 1, 3, 23, 59, 59, 0, time.UTC),
	},
}

func getResult(time time.Time, t *testing.T) *Result {
	t.Helper()
	o := New(time)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	o.Handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	re := &Result{}
	if err := json.Unmarshal(b, &re); err != nil {
		t.Fatal("failed json unmarshal")
	}

	return re
}

func TestOmikuji(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			re := getResult(c.time, t)
			if re.Type != "大吉" {
				t.Errorf("got %v, want %v", re.Type, "大吉")
			}
		})
	}
}
