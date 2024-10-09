package main

import (
	"GoDemo/internal/app"
	"GoDemo/internal/plog"
	"flag"
	"fmt"
	"net/http"
)

const LowestPort = 2_000
const HighestPort = 10_000

func main() {
	portNumber := flag.Int("port", LowestPort, "The port number to listen on")
	flag.Parse()

	application := app.NewApp()

	// Dirty test of the logger:
	application.Logger.Debug("A test debug log...")
	application.Logger.Info("A test info log...")
	application.Logger.Warn("A test warn log...")
	application.Logger.Error("A test error log...")

	foundWorkingSocket := false
	for !foundWorkingSocket {
		application.Logger.Info("Starting server...", plog.KV{Key: "port", Value: *portNumber})
		err := http.ListenAndServe(fmt.Sprintf(":%d", *portNumber), application.Routes())
		if err != nil {
			application.Logger.Error("Failed attempt to connect...", plog.KV{Key: "port", Value: *portNumber})
			*portNumber = *portNumber + 1
		} else if *portNumber > HighestPort {
			application.Logger.Error("Failed to connect to any port.  Stopping.")
			return
		} else {
			foundWorkingSocket = true
			application.Logger.Info("Successfully ran server.", plog.KV{Key: "port", Value: *portNumber})
		}
	}
}
