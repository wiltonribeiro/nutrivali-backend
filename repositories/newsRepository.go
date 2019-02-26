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

//bson.M{ "_id" : 0, "articles" : bson.M{ "$slice": bson.A{"$articles", bson.M{"$subtract": bson.A{bson.M{"$size": "$articles"}, 20}}, bson.M{"$subtract": bson.A{bson.M{"$size": "$articles"}, 20}}}}},

func (repo *NewsRepository) GetNewsArticles(lang string, page int) (articles []models.Article, err error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	end := (10 * page) * -1

	pipeline := []bson.M{
		{"$match": bson.M{"language": lang }},
		{"$project" :
			bson.M{ "_id" : 0, "articles" : bson.M{
				"$cond": bson.M{ "if" : bson.M{ "$gte": bson.A{ bson.M{ "$size": "$articles"}, end * -1 } }, "then" : bson.M { "$slice": bson.A{"$articles", end, 10} }, "else": bson.A{} },
			}},

		},
		{"$lookup": bson.M{
			"from": "articles",
			"localField": "articles",
			"foreignField": "_id",
			"as": "articles",
		}},
	}

	cur, err := config.DB.Collection(repo.Collection).Aggregate(ctx, pipeline)

	if err != nil { return }

	structure := map[string][]models.Article{}

	for cur.Next(ctx) {
		err = cur.Decode(&structure)
		if err != nil { return }
	}


	if err = cur.Err(); err != nil {
		return
	}

	_ = cur.Close(ctx)

	if err != nil {
		return
	}

	articles = structure["articles"]

	return
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

