package webService

import (
	"GoDemo/log"
	"GoDemo/webService/handlers"
	"fmt"
	"net/http"
)

func StartOnPort(portNumber int) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.Ping)
	mux.HandleFunc("/", handlers.Home)
	log.Write("Starting server on port number: %d", portNumber)

	return http.ListenAndServe(fmt.Sprintf(":%d", portNumber), mux)
}
