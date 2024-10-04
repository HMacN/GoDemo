package main

import (
	"net/http"
)

func (app *Application) Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	app.Logger.Info("Ping!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}