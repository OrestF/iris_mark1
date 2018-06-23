package controllers

import "github.com/kataras/iris"

type HomeController struct {
}

func (hc *HomeController) Index(ctx iris.Context) {
	ctx.ViewData("message", "Hellow Go Iris world!")
	ctx.View("home/index.html")
}
