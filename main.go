package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

func main() {
	portNum := 8080
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	logger("Starting server on port number: %d", portNum)

	err := http.ListenAndServe(fmt.Sprintf(":%d", portNum), nil)
	if err != nil {
		logger("Error starting on port number: %d", portNum)
		return
	}
}

func logger(format string, a ...any) {
	log.Println(fmt.Sprintf(format, a...))
}
