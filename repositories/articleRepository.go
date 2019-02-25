package repositories

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"go-app/config"
	"go-app/models"
	"time"
)

type ArticleRepository struct {
	Collection string
}

func (repo *ArticleRepository) GetArticle(id primitive.ObjectID) (article models.Article, err error){


	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"_id": id}
	err = config.DB.Collection(repo.Collection).FindOne(ctx, filter).Decode(&article)

	if err != nil {
		return
	}

	return
}
