package webService

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/ping" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	item := r.URL.Path[1:]
	if item == "" {
		item = "Go"
	}

	switch r.Method {
	case "POST":
		fmt.Fprintf(w, "Hi there, I love %s!", item)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}
}
