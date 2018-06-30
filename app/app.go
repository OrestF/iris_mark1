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

func (*App) Start() {
	app := iris.New()
	configs.ConfigureServer(app)
	configs.ConfigureRoutes(app)
	configs.ConfigureDB()

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
