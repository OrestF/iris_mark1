package configs

import (
	"../../app/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func ConfigureRoutes(app *iris.Application) {
	// ROUTING WAY 1
	homeController := new(controllers.HomeController)

	app.Get("/", func(ctx iris.Context) {
		homeController.Index(ctx)
	}).Name = "root_path"

	// ROUTING WAY 2
	mvc.Configure(app.Party("/users"), controllers.NewUsersController)
	
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	}).Name = "hello_path"
}
