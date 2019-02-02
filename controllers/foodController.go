package controllers

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"go-app/DAOs"
	"go-app/models"
)

type FoodController struct {
	dao DAOs.FoodDAO
}

func (u *FoodController) AddFood(food models.Food) (interface{}, error){
	u.dao = DAOs.FoodDAO{Collection: "foods"}
	food.Id = primitive.NewObjectID()
	return u.dao.AddFood(food)
}

func (u *FoodController) GetFoods() ([]models.Food, error){
	u.dao = DAOs.FoodDAO{Collection: "foods"}
	return u.dao.GetFoods()
}

func (u *FoodController) GetFoodsByDate(time string) ([]models.Food, error){
	u.dao = DAOs.FoodDAO{Collection: "foods"}
	return u.dao.GetFoodsByDate(time)
}

func (u *FoodController) GetFoodsByUser(uid string) ([]models.Food, error){
	u.dao = DAOs.FoodDAO{Collection: "foods"}
	return u.dao.GetFoodsByUser(uid)
}

func (u *FoodController) UpdateFood(food models.Food) error{
	u.dao = DAOs.FoodDAO{Collection: "foods"}
	return u.dao.UpdateFood(food)
}

func (u *FoodController) RemoveFood(food models.Food) error{
	u.dao = DAOs.FoodDAO{Collection: "foods"}
	return u.dao.RemoveFood(food.Id)
}




