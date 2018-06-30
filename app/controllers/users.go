package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"../models"
)

type UsersController struct {
	model models.UserModel
}

func NewUsersController(app *mvc.Application) {
	app.Handle(&UsersController{
		model: models.NewUserModel(),
	})
}

func (uc *UsersController) Get() mvc.Result {
	users := uc.model.All()

	return mvc.View{
		Name: "users/index.html",
		Data: iris.Map{
			"users": users,
		},
	}
}

func (uc *UsersController) GetBy(id int64) mvc.Result {
	user := uc.model.Find(id)

	return mvc.View{
		Name: "users/show.html",
		Data: iris.Map{
			"user": user,
		},
	}
}

func (uc *UsersController) Post(ctx *iris.Context) mvc.Result {
	// email := ctx.URLParam("email")
	params := userParams(*ctx)
	println("VALUES:", params["email"], params["name"])

	user := models.User{Email: params["email"], Name: params["name"]}
	user = uc.model.Create(user)

	// TODO: redirect insted of view render
	return mvc.View{
		Name: "users/show.html",
		Data: iris.Map{
			"user": user,
		},
	}
}

// func (uc *UsersController) PutBy(ctx *iris.Context, id int64) mvc.Result {
// 	params := userParams(*ctx)
// 	println("VALUES:", params["email"], params["name"])

// 	user := models.User{Email: params["email"], Name: params["name"]}
// 	user = uc.model.Create(user)

// 	// TODO: redirect insted of view render
// 	return mvc.View{
// 		Name: "users/show.html",
// 		Data: iris.Map{
// 			"user": user,
// 		},
// 	}
// 	return mvc.View{}
// }

func (uc *UsersController) DeleteBy(id int64) mvc.Result {
	uc.model.Delete(id)

	return mvc.View{
		Name: "users/index.html",
	}
}

// this could be useless if we need different types f.e. int32
func userParams(ctx iris.Context) map[string]string {
	m := map[string]string{
		"email": ctx.FormValue("email"),
		"name":  ctx.FormValue("name"),
	}

	return m
}
