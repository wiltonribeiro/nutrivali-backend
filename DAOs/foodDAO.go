package DAOs

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-app/config"
	"go-app/models"
	"time"
)


type FoodDAO struct {
	config config.Config
	Collection string
}

func (uDAO *FoodDAO) AddFood(food models.Food) (id interface{}, err error){
	uDAO.config = config.Config{}
	db, err := uDAO.config.Connect()
	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := db.Collection(uDAO.Collection).InsertOne(ctx, food)

	if err == nil {id = res.InsertedID}

	defer uDAO.config.Disconnect()
	return
}

func (uDAO *FoodDAO) RemoveFood(food models.Food) (err error){
	uDAO.config = config.Config{}
	db, err := uDAO.config.Connect()
	if err != nil {
		return
	}

	filter := bson.M{"id": food.Id}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	db.Collection(uDAO.Collection).FindOneAndDelete(ctx,filter)

	defer uDAO.config.Disconnect()
	return
}

func (uDAO *FoodDAO) UpdateFood(food models.Food) (err error){
	uDAO.config = config.Config{}
	db, err := uDAO.config.Connect()
	if err != nil {
		return
	}

	filter := bson.M{"_id": food.Id}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	db.Collection(uDAO.Collection).FindOneAndReplace(ctx,filter,food)

	defer uDAO.config.Disconnect()
	return
}

func (uDAO *FoodDAO) GetFoods() (foods []models.Food, err error){
	uDAO.config = config.Config{}
	db, err := uDAO.config.Connect()

	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cur, err := db.Collection(uDAO.Collection).Find(ctx, bson.D{})

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
	defer uDAO.config.Disconnect()
	return
}

func (uDAO *FoodDAO) GetFoodsByUser(userUid string) (foods []models.Food, err error){
	uDAO.config = config.Config{}
	db, err := uDAO.config.Connect()
	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)


	filter := bson.M{"user_uid": userUid}
	cur, err := db.Collection(uDAO.Collection).Find(ctx, filter)
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

	defer uDAO.config.Disconnect()
	return
}



func (uDAO *FoodDAO) GetFoodsByDate(today string, yesterday string) (foods []models.Food, err error){
	uDAO.config = config.Config{}
	db, err := uDAO.config.Connect()
	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//filter := bson.M{"limit_date": t}
	filter := bson.D{
		{"$or",
			bson.A{
				bson.M{"limit_date": today},
				bson.M{"limit_date": yesterday},
			}},
	}

	cur, err := db.Collection(uDAO.Collection).Find(ctx, filter)

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

	defer uDAO.config.Disconnect()
	return
}
