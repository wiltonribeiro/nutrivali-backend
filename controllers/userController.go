package controllers

import (
	"go-app/DAOs"
	"go-app/models"
)

type UserController struct {
	dao DAOs.UserDAO
}

func (u *UserController) AddUser(user models.User) (interface{}, error){
	u.dao = DAOs.UserDAO{Collection: "users"}
	return u.dao.AddUser(user)
}

func (u *UserController) GetUsers() ([]models.User, error) {
	u.dao = DAOs.UserDAO{Collection: "users"}
	return u.dao.GetUsers()
}

func (u *UserController) GetUserById(uid string) (models.User, error){
	u.dao = DAOs.UserDAO{Collection: "users"}
	return u.dao.GetUserById(uid)
}

func (u *UserController) UpdateToken(user models.User) error{
	u.dao = DAOs.UserDAO{Collection: "users"}
	return u.dao.UpdateUserToken(user)
}

func (u *UserController) RemoveUser() {
	//TODO
}

func (u *UserController) UpdateUser() {
	//TODO
}

