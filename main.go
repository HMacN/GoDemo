package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

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

func main() {
	portNum := 8080
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	logger("Starting server on port number: %d", portNum)

	err := http.ListenAndServe(fmt.Sprintf(":%d", portNum), mux)
	if err != nil {
		logger("Error starting on port number: %d", portNum)
		return
	}
}

func logger(format string, a ...any) {
	log.Println(fmt.Sprintf(format, a...))
}
