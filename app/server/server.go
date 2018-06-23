package server

import (
	"github.com/kataras/iris"

	"../../app/configs"
)

type Server struct {
}

func New() (server *Server) {
	return
}

func (s *Server) Start() {
	app := iris.New()
	configs.ConfigureServer(app)
	configs.ConfigureRoutes(app)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
