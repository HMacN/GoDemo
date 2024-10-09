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
		app.TemplateBasePath,
		app.PartialsNavPath,
		app.HomePagePath,
	}

	templateSet, err := template.ParseFiles(filePaths...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = templateSet.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
