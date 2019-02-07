package controllers

import (
	"fmt"
	"github.com/appleboy/go-fcm"
	"go-app/DAOs"
	"go-app/config"
	"go-app/models"
	"time"
)

type NotificationController struct {
	daoFood DAOs.FoodDAO
	daoUser DAOs.UserDAO
}

func (n *NotificationController) RequestNotify() (err error) {
	n.daoFood = DAOs.FoodDAO{Collection: "foods"}
	n.daoUser = DAOs.UserDAO{Collection: "users"}

	dateStrToday := time.Now().Format("02/01/2006")
	dateStrYesterday := time.Now().Add(24*time.Hour).Format("02/01/2006")
	foods, err := n.daoFood.GetFoodsByDate(dateStrToday, dateStrYesterday)

	if err != nil {
		return
	}

	for _, food := range foods {
		user, erro := n.daoUser.GetUserById(food.UserUid)
		if erro != nil {
			return erro
		}

		err = n.notify(food, user.Token)

	}

	return
}

func (n *NotificationController) notify(food models.Food, token string) (err error) {


	c := config.Config{}
	serverKey, _ := c.GetKey()

	msg := &fcm.Message{
		To: token,
		Notification: &fcm.Notification{
			Body: fmt.Sprintf("O seu alimento %v vencerá na data %v. Esteja atento, abraços.", food.Description, food.LimitDate),
			Title: "Olá, preciso te falar algo",
			Sound: "default",
		},
		Data: map[string]interface{}{
			"foo": "bar",
		},
	}

	client, err := fcm.NewClient(serverKey)
	if err != nil {
		return
	}

	_ ,err = client.Send(msg)
	if err != nil {
		return
	}

	return
}
