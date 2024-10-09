package app

import (
	"GoDemo/internal/plog"
	"path/filepath"
	"runtime"
	"strings"
)

const TemplateBaseFilePath = "..\\..\\ui\\html\\base.html"
const PartialsNavFilePath = "..\\..\\ui\\html\\partials\\nav.html"
const HomePageFilePath = "..\\..\\ui\\html\\pages\\home.html"
const StaticFilePath = "..\\..\\ui\\static\\"

type Application struct {
	Logger           plog.LogWrapper
	TemplateBasePath string
	PartialsNavPath  string
	HomePagePath     string
	StaticPath       string
}

func NewApp() Application {
	var (
		_, callingFile, _, _ = runtime.Caller(0)
		appPath              = filepath.Dir(callingFile)
	)

	return Application{
		Logger:           plog.New(),
		TemplateBasePath: strings.Join([]string{appPath, TemplateBaseFilePath}, "\\"),
		PartialsNavPath:  strings.Join([]string{appPath, PartialsNavFilePath}, "\\"),
		HomePagePath:     strings.Join([]string{appPath, HomePageFilePath}, "\\"),
		StaticPath:       strings.Join([]string{appPath, StaticFilePath}, "\\"),
	}
}
