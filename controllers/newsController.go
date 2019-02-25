package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"go-app/models"
	"go-app/repositories"
	"strconv"
)

type NewsController struct {
	repo repositories.NewsRepository
}

func (u *NewsController) GetNews(ctx iris.Context) {

	lang := ctx.Params().Get("lang")
	page := ctx.Params().Get("page")

	u.repo = repositories.NewsRepository{Collection: "news"}
	artRepo := repositories.ArticleRepository{Collection: "articles"}

	var articles []models.Article
	news , err := u.repo.GetNews(lang)


	p, err := strconv.Atoi(page)

	begin := (len(news.Articles) - 1) - (15 * p) + 15
	end := begin - 15

	if end < 0 {
		end = 0
	}

	if begin > 0 {
		for i:=begin; i > end; i--{

			value := news.Articles[i]
			article, err := artRepo.GetArticle(value)

			if err != nil {
				break
			}

			articles = append(articles, article)

			if len(articles) == 15 {
				break
			}

		}
	} else {
		ctx.StatusCode(404)
	}


	if err != nil{
		fmt.Println(err.Error())
		ctx.StatusCode(500)
	} else{
		_ , _ = ctx.JSON(articles)
	}
}