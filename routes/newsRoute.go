package routes

import (
	"github.com/kataras/iris"
	"go-app/controllers"
	"go-app/models"
)

var NewsRoute = models.Route{
	Apply: func(app *iris.Application) {

		var newsController = controllers.NewsController{}

		app.Handle("GET", "news/{lang}/{page}", newsController.GetNews)

	},
}
