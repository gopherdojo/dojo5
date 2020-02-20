package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gopherdojo/dojo5/kadai4/ebiiim/pkg/omikuji"
)

func requireUserAgent(h http.HandlerFunc, ua string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		isCurl := strings.Contains(r.UserAgent(), ua)
		if !isCurl {
			http.Error(w, fmt.Sprintf("Please use %s.", ua), http.StatusForbidden)
			return
		}
		h(w, r)
	}
	return http.HandlerFunc(fn)
}

func main() {
	k := omikuji.NewSimpleOmikuji(rand.New(rand.NewSource(time.Now().UnixNano())))
	s := omikuji.NewKujiServer(k)
	mux := http.NewServeMux()
	mux.HandleFunc("/", requireUserAgent(s.TmplNamedKuji, "curl"))
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start the server %v", err)
	}
}
