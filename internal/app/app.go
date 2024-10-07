package app

import (
	"GoDemo/internal/logWrapper"
	"path/filepath"
	"runtime"
	"strings"
)

const TemplateBaseFilePath = "..\\..\\ui\\html\\base.html"
const PartialsNavFilePath = "..\\..\\ui\\html\\partials\\nav.html"
const HomePageFilePath = "..\\..\\ui\\html\\pages\\home.html"

type Application struct {
	Logger           logWrapper.LogWrapper
	TemplateBasePath string
	PartialsNavPath  string
	HomePagePath     string
}

func NewApp() Application {
	var (
		_, callingFile, _, _ = runtime.Caller(0)
		appPath              = filepath.Dir(callingFile)
	)

	return Application{
		Logger:           logWrapper.New(),
		TemplateBasePath: strings.Join([]string{appPath, TemplateBaseFilePath}, "\\"),
		PartialsNavPath:  strings.Join([]string{appPath, PartialsNavFilePath}, "\\"),
		HomePagePath:     strings.Join([]string{appPath, HomePageFilePath}, "\\"),
	}
}
