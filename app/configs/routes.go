package configs

import (
	"../../app/controllers"
	"github.com/kataras/iris"
)

func ConfigureRoutes(app *iris.Application) {
	homeController := controllers.NewHomeController()

	app.Get("/", func(ctx iris.Context) {
		homeController.Index(ctx)
	})

	app.Get("/cars", func(ctx iris.Context) {
		homeController.Cars(ctx)
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
}
