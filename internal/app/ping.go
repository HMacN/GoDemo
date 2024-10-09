package app

import (
	"net/http"
)

func (app *Application) Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	app.Logger.Info("Ping!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
