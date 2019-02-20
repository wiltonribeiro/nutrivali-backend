package repositories

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-app/config"
	"go-app/models"
	"time"
)


type FoodRepository struct {
	Collection string
}

func (repo *FoodRepository) AddFood(food models.Food) (id interface{}, err error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := config.DB.Collection(repo.Collection).InsertOne(ctx, food)

	if err == nil {id = res.InsertedID}

	return
}

func (repo *FoodRepository) RemoveFood(food models.Food) (err error){

	filter := bson.M{"id": food.Id}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	config.DB.Collection(repo.Collection).FindOneAndDelete(ctx,filter)

	return
}

func (repo *FoodRepository) UpdateFood(food models.Food) (err error){

	filter := bson.M{"_id": food.Id}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	config.DB.Collection(repo.Collection).FindOneAndReplace(ctx,filter,food)

	return
}

func (repo *FoodRepository) GetFoods() (foods []models.Food, err error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := config.DB.Collection(repo.Collection).Find(ctx, bson.D{})

	if err != nil { return }

	for cur.Next(ctx) {
		var food models.Food
		err = cur.Decode(&food)
		if err != nil { return }
		foods = append(foods, food)
	}

	if err = cur.Err(); err != nil {
		return
	}

	_ = cur.Close(ctx)

	return
}

func (repo *FoodRepository) GetFoodsByUser(userUid string) (foods []models.Food, err error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)


	filter := bson.M{"user_uid": userUid}
	cur, err := config.DB.Collection(repo.Collection).Find(ctx, filter)
	if err != nil { return }

	for cur.Next(ctx) {
		var food models.Food
		err = cur.Decode(&food)
		if err != nil { return }
		foods = append(foods, food)
	}

	if err = cur.Err(); err != nil {
		return
	}

	_ = cur.Close(ctx)


	return
}


func (repo *FoodRepository) GetFoodsByDate(today string, yesterday string) (foods []models.Food, err error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//filter := bson.M{"limit_date": t}
	filter := bson.D{
		{"$or",
			bson.A{
				bson.M{"limit_date": today},
				bson.M{"limit_date": yesterday},
			}},
	}

	cur, err := config.DB.Collection(repo.Collection).Find(ctx, filter)

	if err != nil { return }

	for cur.Next(ctx) {
		var food models.Food
		err = cur.Decode(&food)
		if err != nil { return }
		foods = append(foods, food)
	}

	if err = cur.Err(); err != nil {
		return
	}

	_ = cur.Close(ctx)

	if err != nil {
		return
	}

	return
}
