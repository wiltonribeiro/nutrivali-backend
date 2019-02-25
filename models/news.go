package models

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type News struct {
	Language string `bson:"language" json:"language"`
	Articles []primitive.ObjectID `bson:"articles" json:"articles"`
}
