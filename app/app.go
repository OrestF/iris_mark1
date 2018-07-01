package app

import (
	"github.com/kataras/iris"

	"../app/configs"
)

type App struct {
}

func New() (app *App) {
	return
}

func (app *App) Start() {
	irisApp := iris.New()
	configs.ConfigureServer(irisApp)
	configs.ConfigureRoutes(irisApp)
	configs.ConfigureDB()

	irisApp.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
