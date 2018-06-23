package controllers

import "github.com/kataras/iris"

type HomeController struct {
}

func NewHomeController() HomeController {
	return HomeController{}
}

func (hc HomeController) Index(ctx iris.Context) {
	ctx.View("home/index.html")
}

func (hc HomeController) Cars(ctx iris.Context) {
	ctx.HTML("Cars list")
}
