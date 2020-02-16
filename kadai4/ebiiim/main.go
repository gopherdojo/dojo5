package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gopherdojo/dojo5/kadai4/ebiiim/pkg/omikuji"
)

func main() {
	kuji := omikuji.NewSimpleOmikuji(rand.New(rand.NewSource(time.Now().UnixNano())))
	kujiServer := omikuji.NewKujiServer(kuji)
	mux := http.NewServeMux()
	mux.HandleFunc("/", kujiServer.Handler)
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start the server %v", err)
	}
}
