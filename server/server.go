package server

import (
	"github.com/kataras/iris"
	"go-app/models"
	"go-app/routes"
)

func applyRoutes(routes []models.Route){

	app := iris.Default()

	for _,item := range routes {
		item.Apply(app)
	}

	app.Run(iris.Addr(""))
}

func InitServer(){
	var r = []models.Route {
		routes.NewsRoute,
		routes.UserRouter,
		routes.FoodRoute,
		routes.NotificationRoute,
	}
	applyRoutes(r)
}
