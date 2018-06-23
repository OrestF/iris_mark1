package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type UsersController struct {
}

func NewUsersController(app *mvc.Application) {
	app.Handle(new(UsersController))
}

func (uc *UsersController) Get() mvc.Result {
	return mvc.View{
		Name: "users/index.html",
	}
}

func (uc *UsersController) GetBy(id string) mvc.Result {
	return mvc.View{
		Name: "users/show.html",
		Data: iris.Map{
			"id": id,
		},
	}
}
