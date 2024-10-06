package main

import (
	"GoDemo/internal/app"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	portNumber := flag.Int("port", 8080, "The port number to listen on")
	flag.Parse()

	app := app.NewApp()
	app.Logger.Info("Attempting server start on port %d", portNumber)
	//server := &http.Server{
	//	Addr: fmt.Sprintf(":%d", portNumber),
	//}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", app.Ping)
	mux.HandleFunc("/snippet", app.Snippet)
	mux.HandleFunc("/", app.Home)
	app.Logger.Info("Starting server on port number: %d", portNumber)

	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), mux)
	if err != nil {
		app.Logger.Error("Error starting on port number: %d", portNumber)
		return
	}
}
