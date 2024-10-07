package app

import (
	"net/http"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

}
