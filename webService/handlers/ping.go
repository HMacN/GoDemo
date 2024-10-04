package handlers

import "net/http"

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ping" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
		return
	}
}
