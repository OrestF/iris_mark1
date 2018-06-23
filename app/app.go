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

func (a *App) Start() {
	app := iris.New()
	configs.ConfigureServer(app)
	configs.ConfigureRoutes(app)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
