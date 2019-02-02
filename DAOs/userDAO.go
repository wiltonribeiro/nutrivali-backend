package DAOs

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-app/config"
	"go-app/models"
	"time"
)


type UserDAO struct {
	config config.Config
	Collection string
}

func (uDAO *UserDAO) AddUser(user models.User) (id interface{}, err error){
	uDAO.config = config.Config{}
	db, err := uDAO.config.Connect()
	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := db.Collection(uDAO.Collection).InsertOne(ctx, user)
	id = res.InsertedID

	uDAO.config.Disconnect()
	return
}

func (uDAO *UserDAO) GetUsers() (users []models.User, err error){
	uDAO.config = config.Config{}
	db, err := uDAO.config.Connect()
	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)


	cur, err := db.Collection(uDAO.Collection).Find(ctx, bson.D{})
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

	uDAO.config.Disconnect()
	return
}

func (uDAO *UserDAO) GetUserById(uid string) (user models.User, err error){
	uDAO.config = config.Config{}
	db, err := uDAO.config.Connect()
	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)


	filter := bson.M{"auth_uid": uid}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = db.Collection(uDAO.Collection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return
	}

	uDAO.config.Disconnect()
	return
}






