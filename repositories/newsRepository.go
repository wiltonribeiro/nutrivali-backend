package repositories

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-app/config"
	"go-app/models"
	"time"
)

type NewsRepository struct {
	Collection string
}

func (repo *NewsRepository) GetNews(lang string) (news models.News, err error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{"language": lang}
	cur, err := config.DB.Collection(repo.Collection).Find(ctx, filter)

	if err != nil {
		return
	}

	for cur.Next(ctx) {
		err = cur.Decode(&news)
		if err != nil {
			return
		}

		return
	}

	if err = cur.Err(); err != nil {
		return
	}

	_ = cur.Close(ctx)


	return
}

