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

	filePaths := []string{
		TemplateBaseFilePath,
		HomePageFilePath,
	}

	templateSet, err := template.ParseFiles(filePaths...)
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = templateSet.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
