package app

import "net/http"

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(app.StaticPath))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/ping", app.Ping)
	mux.HandleFunc("/snippet", app.Snippet)
	mux.HandleFunc("/", app.Home)
	return mux
}
