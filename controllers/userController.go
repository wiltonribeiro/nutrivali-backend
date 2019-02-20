package controllers

import (
	"github.com/kataras/iris"
	"go-app/repositories"
	"go-app/models"
)

type UserController struct {
	dao repositories.UserRepository
}

func (u *UserController) AddUser(ctx iris.Context) {

	var user models.User
	u.dao = repositories.UserRepository{Collection: "users"}

	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(400)
	} else {
		_, err := u.dao.AddUser(user)
		if err != nil {
			ctx.StatusCode(500)
		} else {
			ctx.StatusCode(200)
		}
	}

}

func (u *UserController) GetUsers(ctx iris.Context)  {

	u.dao = repositories.UserRepository{Collection: "users"}
	users, err := u.dao.GetUsers()

	if err != nil && err.Error() == "mongo: no documents in result" {
		ctx.StatusCode(404)
	} else if err != nil {
		ctx.StatusCode(500)
	} else {
		_, _ = ctx.JSON(users)
	}

}


func (u *UserController) UpdateToken(ctx iris.Context) {
	u.dao = repositories.UserRepository{Collection: "users"}

	var user models.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(400)
	} else {
		err := u.dao.UpdateUserToken(user)
		if err != nil {
			ctx.StatusCode(500)
		} else {
			ctx.StatusCode(200)
		}
	}

}


func (u *UserController) GetUserById(ctx iris.Context) {
	u.dao = repositories.UserRepository{Collection: "users"}

	userUid := ctx.Params().Get("uid")
	user , err := u.dao.GetUserById(userUid)

	if err != nil && err.Error() == "mongo: no documents in result" {
		ctx.StatusCode(404)
	} else if err != nil {
		ctx.StatusCode(500)
	} else {
		_, _ = ctx.JSON(user)
	}

}


func (u *UserController) RemoveUser() {
	//TODO
}

