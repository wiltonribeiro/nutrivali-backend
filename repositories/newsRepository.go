package repositories

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-app/config"
	"go-app/models"
	"sort"
	"time"
)

type NewsRepository struct {
	Collection string
}

func (repo *NewsRepository) GetNewsArticles(lang string, page int) (articles []models.Article, err error){

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	count := 30

	end := (count * page) * -1

	pipeline := []bson.M{
		{"$match": bson.M{"language": lang }},
		{"$project" :
			bson.M{ "_id" : 0, "articles" : bson.M{
				"$cond": bson.M{ "if" : bson.M{ "$gte": bson.A{ bson.M{ "$size": "$articles"}, end * -1 } }, "then" : bson.M { "$slice": bson.A{"$articles", end, count} }, "else": bson.A{} },
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

	articles = sortArticles(structure["articles"])

	return
}

func sortArticles(articles []models.Article) []models.Article{
	sort.SliceStable(articles, func(i, j int) bool {

		ai, aj := articles[i], articles[j]


		t1, _ := time.Parse(time.RFC3339Nano, ai.PublishedAt)
		t2, _ := time.Parse(time.RFC3339Nano, aj.PublishedAt)

		switch {
		case t1.After(t2):
			return true
		default:
			return false
		}
	})

	return articles
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

