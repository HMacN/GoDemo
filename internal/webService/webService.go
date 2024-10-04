package webService

import (
	"GoDemo/internal/log"
	handlers2 "GoDemo/internal/webService/handlers"
	"fmt"
	"net/http"
)

func StartOnPort(portNumber int) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers2.Ping)
	mux.HandleFunc("/snippet", handlers2.Snippet)
	mux.HandleFunc("/", handlers2.Home)
	log.Write("Starting server on port number: %d", portNumber)

	return http.ListenAndServe(fmt.Sprintf(":%d", portNumber), mux)
}
