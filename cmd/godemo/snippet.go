package main

import (
	"net/http"
	"strconv"
)

func (app *Application) Snippet(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	app.Logger.Info("Testing snippet system.")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(id)))
}
