package repositories

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-app/config"
	"go-app/models"
	"time"
)


type UserRepository struct {
	Collection string
}

func (repo *UserRepository) AddUser(user models.User) (id interface{}, err error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := config.DB.Collection(repo.Collection).InsertOne(ctx, user)
	id = res.InsertedID

	return
}

func (repo *UserRepository) GetUsers() (users []models.User, err error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := config.DB.Collection(repo.Collection).Find(ctx, bson.D{})

	if err != nil { return }

	for cur.Next(ctx) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil { return }
		users = append(users, user)
	}

	if err = cur.Err(); err != nil {
		return
	}

	_ = cur.Close(ctx)

	return
}

func (repo *UserRepository) GetUserById(uid string) (user models.User, err error){

	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)


	filter := bson.M{"auth_uid": uid}

	err = config.DB.Collection(repo.Collection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return
	}

	return
}

func (repo *UserRepository) UpdateUserToken(user models.User) (err error){


	filter := bson.M{"auth_uid": user.AuthUid}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var ai = config.DB.Collection(repo.Collection).FindOneAndReplace(ctx,filter,user)
	if ai.Err() != nil {
		return ai.Err()
	}

	return
}






