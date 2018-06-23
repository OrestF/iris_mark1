package configs

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func ConfigureServer(app *iris.Application) {
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	registerViews(app)
}

func registerViews(app *iris.Application) {
	app.RegisterView(iris.HTML("./app/views", ".html"))
}
