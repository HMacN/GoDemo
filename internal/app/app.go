package app

import "GoDemo/internal/logWrapper"

const TemplateBaseFilePath = "../../ui/html/base.html"
const HomePageFilePath = "../../ui/html/pages/home.html"

type Application struct {
	Logger logWrapper.LogWrapper
}

func NewApp() Application {
	return Application{Logger: logWrapper.New()}
}
