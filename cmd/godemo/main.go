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

	app.Logger.Info("Starting server on port number: %d", portNumber)

	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), app.Routes())
	if err != nil {
		app.Logger.Error("Error starting on port number: %d", portNumber)
		return
	}
}
