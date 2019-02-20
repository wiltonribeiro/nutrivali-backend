package routes

import (
	"github.com/kataras/iris"
	"go-app/controllers"
	"go-app/models"
)

var FoodRoute = models.Route{
	Apply: func(app *iris.Application) {

		var foodController = controllers.FoodController{}

		app.Handle("GET", "foods", foodController.GetFoods)

		app.Handle("POST", "foods/date", foodController.GetFoodsByDate)

		app.Handle("GET", "foods/user/{uid}", foodController.GetFoodsByUser)

		app.Handle("POST", "foods", foodController.AddFood)

		app.Handle("POST", "foods/update", foodController.UpdateFood)

		app.Handle("POST", "foods/remove", foodController.RemoveFood)
	},
}
