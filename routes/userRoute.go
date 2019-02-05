package routes

import (
	"github.com/kataras/iris"
	"go-app/controllers"
	"go-app/models"
)

var UserRouter = models.Route{
	Apply: func(app *iris.Application) {

		var userController = controllers.UserController{}

		app.Handle("GET", "/", func(ctx iris.Context) {
			_, _ = ctx.HTML("<p>TUDO OK</p>")
		})

		app.Handle("GET", "/users", func(ctx iris.Context) {
			users , err := userController.GetUsers()

			if err != nil && err.Error() == "mongo: no documents in result" {
				ctx.StatusCode(404)
			} else if err != nil {
				ctx.StatusCode(500)
			} else {
				_, _ = ctx.JSON(users)
			}
		})


		app.Handle("POST", "/users", func(ctx iris.Context) {

			var user models.User
			err := ctx.ReadJSON(&user)
			if err != nil {
				ctx.StatusCode(400)
			} else {
				_, err := userController.AddUser(user)
				if err != nil {
					ctx.StatusCode(500)
				} else {
					ctx.StatusCode(200)
				}
			}
		})

		app.Handle("POST", "/users/update", func(ctx iris.Context) {

			var user models.User
			err := ctx.ReadJSON(&user)
			if err != nil {
				ctx.StatusCode(400)
			} else {
				err := userController.UpdateToken(user)
				if err != nil {
					ctx.StatusCode(500)
				} else {
					ctx.StatusCode(200)
				}
			}
		})

		app.Handle("GET", "/user/{uid}", func(ctx iris.Context) {
			userUid := ctx.Params().Get("uid")
			user , err := userController.GetUserById(userUid)

			if err != nil && err.Error() == "mongo: no documents in result" {
				ctx.StatusCode(404)
			} else if err != nil {
				ctx.StatusCode(500)
			} else {
				_, _ = ctx.JSON(user)
			}
		})
	},
}


