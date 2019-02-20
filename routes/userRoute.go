package routes

import (
	"github.com/kataras/iris"
	"go-app/controllers"
	"go-app/models"
	"go-app/utils"
)

var UserRouter = models.Route{
	Apply: func(app *iris.Application) {

		var userController = controllers.UserController{}

		app.Handle("GET", "/", func(ctx iris.Context) {
			_,_ = ctx.JSON(utils.GetLog())
		})

		app.Handle("GET", "/users", userController.GetUsers)

		app.Handle("POST", "/users", userController.AddUser)

		app.Handle("POST", "/users/update", userController.UpdateToken)

		app.Handle("GET", "/user/{uid}", userController.GetUserById)

	},
}


