package routes

import (
	"github.com/kataras/iris"
	"go-app/controllers"
	"go-app/models"
)

var NotificationRoute = models.Route{
	Apply: func(app *iris.Application) {

		var notificationController = controllers.NotificationController{}

		app.Handle("GET", "/notify", notificationController.Notify)

	},
}