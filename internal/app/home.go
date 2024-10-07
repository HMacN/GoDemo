package app

import (
	"html/template"
	"net/http"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	template, err := template.ParseFiles(HomePageFilePath)
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, nil)
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
