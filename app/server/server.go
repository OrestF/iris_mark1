package server

import (
	"github.com/kataras/iris"

	"../../app/controllers"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

type Server struct {
}

func New() (server *Server) {
	return
}

func (s *Server) Start() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	//VIEWS
	app.RegisterView(iris.HTML("./app/views", ".html"))

	// CONTROLLERS
	homeController := controllers.NewHomeController()

	// Method:   GET
	// Resource: http://localhost:8080
	// app.Handle("GET", "/", func(ctx iris.Context) {
	// 	ctx.HTML("<h1>Welcome</h1>")
	// })

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

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
