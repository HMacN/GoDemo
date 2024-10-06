package app

import "net/http"

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", app.Ping)
	mux.HandleFunc("/snippet", app.Snippet)
	mux.HandleFunc("/", app.Home)
	return mux
}
