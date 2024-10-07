package app

import "GoDemo/internal/logWrapper"

const HomePageFilePath = "../../ui/html/pages/home.html"

type Application struct {
	Logger logWrapper.LogWrapper
}

func NewApp() Application {
	return Application{Logger: logWrapper.New()}
}
