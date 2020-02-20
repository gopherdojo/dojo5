package omikuji

import (
	"fmt"
	"html"
	"net/http"
)

type KujiServer struct {
	kuji lottery
}

func NewKujiServer(kuji lottery) *KujiServer {
	ks := &KujiServer{}
	ks.kuji = kuji
	return ks
}

func (ks *KujiServer) SimpleKuji(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, ks.kuji.Do())
	if err != nil {
		// TODO: handles error
	}
}

func (ks *KujiServer) NamedKuji(w http.ResponseWriter, r *http.Request) {
	name := html.EscapeString(r.FormValue("name"))
	if name == "" {
		name = "名無し"
	}
	_, err := fmt.Fprintf(w, "%sさんの本日の運勢は「%s」です！", name, ks.kuji.Do())
	if err != nil {
		// TODO: handles error
	}
}
