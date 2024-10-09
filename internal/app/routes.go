package app

import (
	"GoDemo/internal/utils"
	"net/http"
)

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileSystem := utils.SafeFileSystem{Files: http.Dir(app.StaticPath)}
	fileServer := http.FileServer(fileSystem)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/ping", app.Ping)
	mux.HandleFunc("/snippet", app.Snippet)
	mux.HandleFunc("/", app.Home)
	return mux
}
