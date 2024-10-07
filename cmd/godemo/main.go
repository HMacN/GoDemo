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
	app.Logger.Info("Attempting server start on port %d", *portNumber)
	//server := &http.Server{
	//	Addr: fmt.Sprintf(":%d", portNumber),
	//}

	app.Logger.Info("Starting server on port number: %d", *portNumber)

	foundWorkingSocket := false
	for !foundWorkingSocket {
		err := http.ListenAndServe(fmt.Sprintf(":%d", *portNumber), app.Routes())
		if err != nil {
			app.Logger.Error("Failed attempt to connect on port number: %d", portNumber)
			*portNumber = *portNumber + 1
		} else if !foundWorkingSocket && *portNumber > 9000 {
			app.Logger.Error("Failed to connect to any port.  Stopping.")
			return
		} else {
			foundWorkingSocket = true
			app.Logger.Info("Successfully ran server on port number: %d", *portNumber)
		}
	}
}
