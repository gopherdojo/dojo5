package omikuji

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Omikuji contains time
type Omikuji struct {
	time time.Time
}

// Result contains type of omikuji result
type Result struct {
	Type string `json:"type"`
}

var types = []string{
	"大吉",
	"吉",
	"中吉",
	"小吉",
	"末吉",
	"凶",
	"大凶",
}

// New returns Omikuji instance
func New(t time.Time) *Omikuji {
	rand.Seed(t.UnixNano())
	return &Omikuji{time: t}
}

// Handler returns omikuji result
func (o *Omikuji) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	result := draw(o.time)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println("Error:", err)
	}
}

func draw(t time.Time) *Result {
	yd := t.YearDay()

	if yd == 1 || yd == 2 || yd == 3 {
		return &Result{Type: types[0]}
	}

	return &Result{Type: types[rand.Intn(len(types))]}
}
