package app

import "GoDemo/internal/logWrapper"

type Application struct {
	Logger logWrapper.LogWrapper
}

func NewApp() Application {
	return Application{Logger: logWrapper.New()}
}
