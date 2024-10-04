package webService

import (
	"GoDemo/log"
	"fmt"
	"net/http"
)

func StartOnPort(portNumber int) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Write("Starting server on port number: %d", portNumber)

	return http.ListenAndServe(fmt.Sprintf(":%d", portNumber), mux)
}
