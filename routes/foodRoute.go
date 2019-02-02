package routes

import (
	"github.com/kataras/iris"
	"go-app/controllers"
	"go-app/models"
)

var FoodRoute = models.Route{
	Apply: func(app *iris.Application) {

		var foodController = controllers.FoodController{}

		app.Handle("GET", "foods", func(ctx iris.Context) {

			foods, err := foodController.GetFoods()

			if (err != nil && err.Error() == "mongo: no documents in result") || len(foods) == 0 {
				ctx.StatusCode(404)
			} else if err != nil {
				ctx.StatusCode(500)
			} else {
				_, _ = ctx.JSON(foods)
			}
		})

		app.Handle("GET", "foods/date/{date}", func(ctx iris.Context) {

			foods, err := foodController.GetFoodsByDate(ctx.Params().Get("date"))

			if (err != nil && err.Error() == "mongo: no documents in result") || len(foods) == 0 {
				ctx.StatusCode(404)
			} else if err != nil {
				ctx.StatusCode(500)
			} else {
				_, _ = ctx.JSON(foods)
			}
		})

		app.Handle("GET", "foods/user/{uid}", func(ctx iris.Context) {

			foods, err := foodController.GetFoodsByUser(ctx.Params().Get("uid"))

			if (err != nil && err.Error() == "mongo: no documents in result") || len(foods) == 0 {
				ctx.StatusCode(404)
			} else if err != nil {
				ctx.StatusCode(500)
			} else {
				_, _ = ctx.JSON(foods)
			}
		})

		app.Handle("POST", "foods", func(ctx iris.Context) {
			var food models.Food

			err := ctx.ReadJSON(&food)
			if err != nil {
				ctx.StatusCode(400)
			} else {
				_, err := foodController.AddFood(food)
				if err != nil {
					ctx.StatusCode(500)
				} else {
					ctx.StatusCode(200)
				}
			}
		})

		app.Handle("POST", "foods/update", func(ctx iris.Context) {
			var food models.Food

			err := ctx.ReadJSON(&food)
			if err != nil {
				ctx.StatusCode(400)
			} else {
				err := foodController.UpdateFood(food)
				if err != nil {
					ctx.StatusCode(500)
				} else {
					ctx.StatusCode(200)
				}
			}
		})

		app.Handle("POST", "foods/remove", func(ctx iris.Context) {
			var food models.Food
			err := ctx.ReadJSON(&food)

			if err != nil {
				ctx.StatusCode(400)
			} else {
				err := foodController.RemoveFood(food)
				if err != nil {
					ctx.StatusCode(500)
				} else {
					ctx.StatusCode(200)
				}
			}
		})
	},
}
