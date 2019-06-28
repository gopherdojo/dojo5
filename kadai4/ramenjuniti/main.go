package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gopherdojo/dojo5/kadai4/ramenjuniti/omikuji"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8080, "port number")
}

func main() {
	flag.Parse()
	o := omikuji.New(time.Now())
	http.HandleFunc("/", o.Handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
