package controllers

import (
	"github.com/kataras/iris"
	"go-app/repositories"
	"go-app/models"
)

type FoodController struct {
	dao repositories.FoodRepository
}


func (u *FoodController) GetFoods(ctx iris.Context) {
	u.dao = repositories.FoodRepository{Collection: "foods"}
	foods, err := u.dao.GetFoods()

	if (err != nil && err.Error() == "mongo: no documents in result") || len(foods) == 0 {
		ctx.StatusCode(404)
	} else if err != nil {
		ctx.StatusCode(500)
	} else {
		_, _ = ctx.JSON(foods)
	}

}


func (u *FoodController) GetFoodsByDate(ctx iris.Context) {
	u.dao = repositories.FoodRepository{Collection: "foods"}

	var value = struct {
		Date string `json:"date"`
	}{}

	_ = ctx.ReadJSON(&value)

	var time = value.Date
	foods, err := u.dao.GetFoodsByDate(time, time)
	if err != nil {
		print(err.Error())
	}

	if (err != nil && err.Error() == "mongo: no documents in result") || len(foods) == 0 {
		ctx.StatusCode(404)
	} else if err != nil {
		ctx.StatusCode(500)
	} else {
		_, _ = ctx.JSON(foods)
	}

}


func (u *FoodController) GetFoodsByUser(ctx iris.Context) {
	u.dao = repositories.FoodRepository{Collection: "foods"}

	var uid = ctx.Params().Get("uid")
	foods, err := u.dao.GetFoodsByUser(uid)

	if (err != nil && err.Error() == "mongo: no documents in result") || len(foods) == 0 {
		ctx.StatusCode(404)
	} else if err != nil {
		ctx.StatusCode(500)
	} else {
		_, _ = ctx.JSON(foods)
	}

}


func (u *FoodController) AddFood(ctx iris.Context) {
	u.dao = repositories.FoodRepository{Collection: "foods"}

	var food models.Food
	err := ctx.ReadJSON(&food)

	if err != nil {
		ctx.StatusCode(400)
	} else {
		_, err := u.dao.AddFood(food)
		if err != nil {
			ctx.StatusCode(500)
		} else {
			ctx.StatusCode(200)
		}
	}

}

func (u *FoodController) UpdateFood(ctx iris.Context) {
	u.dao = repositories.FoodRepository{Collection: "foods"}

	var food models.Food
	err := ctx.ReadJSON(&food)

	if err != nil {
		ctx.StatusCode(400)
	} else {
		err := u.dao.UpdateFood(food)
		if err != nil {
			ctx.StatusCode(500)
		} else {
			ctx.StatusCode(200)
		}
	}

}

func (u *FoodController) RemoveFood(ctx iris.Context) {
	u.dao = repositories.FoodRepository{Collection: "foods"}

	var food models.Food
	err := ctx.ReadJSON(&food)

	if err != nil {
		ctx.StatusCode(400)
	} else {
		err := u.dao.RemoveFood(food)
		if err != nil {
			ctx.StatusCode(500)
		} else {
			ctx.StatusCode(200)
		}
	}
}




