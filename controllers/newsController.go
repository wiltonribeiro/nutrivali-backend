package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"go-app/repositories"
	"strconv"
)

type NewsController struct {
	repo repositories.NewsRepository
}

func (u *NewsController) GetNewsArticles(ctx iris.Context) {

	u.repo = repositories.NewsRepository{Collection: "news"}

	lang := ctx.Params().Get("lang")
	p := ctx.Params().Get("page")

	page, err := strconv.Atoi(p)

	if err != nil { ctx.StatusCode(400) }

	articles, err := u.repo.GetNewsArticles(lang, page)

	if err != nil {
		fmt.Println(err.Error())
		ctx.StatusCode(500)
	} else{
		_ , _ = ctx.JSON(articles)
	}

}