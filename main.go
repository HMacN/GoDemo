package main

import (
	"GoDemo/log"
	"GoDemo/webService"
)

func main() {
	portNumber := 8080
	err := webService.StartOnPort(portNumber)
	if err != nil {
		log.Write("Error starting on port number: %d", portNumber)
		return
	}
}
