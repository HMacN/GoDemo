package app

import (
	"GoDemo/internal/plog"
	"database/sql"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

const TemplateBaseFilePath = "ui\\html\\base.html"
const PartialsNavFilePath = "ui\\html\\partials\\nav.html"
const HomePageFilePath = "ui\\html\\pages\\home.html"
const StaticFilePath = "ui\\static\\"
const PathToProjectRoot = "..\\..\\..\\"
const PathToLogFile = "\\logs"

type Application struct {
	Logger           plog.LogWrapper
	Database         *sql.DB
	TemplateBasePath string
	PartialsNavPath  string
	HomePagePath     string
	StaticPath       string
}

func NewApp() Application {
	var (
		_, callingFile, _, _ = runtime.Caller(0)
		thisFilePath         = filepath.Dir(callingFile)
	)

	rootPath := filepath.Clean(thisFilePath + PathToProjectRoot)

	return Application{
		Logger:           plog.New(rootPath+PathToLogFile, "myLog", "log"),
		Database:         nil,
		TemplateBasePath: strings.Join([]string{rootPath, TemplateBaseFilePath}, "\\"),
		PartialsNavPath:  strings.Join([]string{rootPath, PartialsNavFilePath}, "\\"),
		HomePagePath:     strings.Join([]string{rootPath, HomePageFilePath}, "\\"),
		StaticPath:       strings.Join([]string{rootPath, StaticFilePath}, "\\"),
	}
}

func (app *Application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.Logger.Error(
		err.Error(),
		plog.KV{Key: "method", Value: method},
		plog.KV{Key: "uri", Value: uri})
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
