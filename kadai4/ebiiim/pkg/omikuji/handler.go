package omikuji

import (
	"fmt"
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

func (ks *KujiServer) Handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, ks.kuji.Do())
	if err != nil {
		// TODO: handles error
	}
}
