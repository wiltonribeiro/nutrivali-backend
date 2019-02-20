package services

import (
	"fmt"
	"github.com/appleboy/go-fcm"
	"go-app/repositories"
	"go-app/config"
	"go-app/models"
	"time"
)

type NotificationService struct {
	daoFood repositories.FoodRepository
	daoUser repositories.UserRepository
}

func (n *NotificationService) RequestNotify() (err error) {
	n.daoFood = repositories.FoodRepository{Collection: "foods"}
	n.daoUser = repositories.UserRepository{Collection: "users"}

	dateStrToday := time.Now().Format("02/01/2006")
	dateStrYesterday := time.Now().Add(24*time.Hour).Format("02/01/2006")
	foods, err := n.daoFood.GetFoodsByDate(dateStrToday, dateStrYesterday)

	if err != nil {
		return
	}

	for _, food := range foods {
		user, err1 := n.daoUser.GetUserById(food.UserUid)
		if err1 != nil {
			return err1
		}

		err = n.notify(food, user)

	}

	return
}

func (n *NotificationService) notify(food models.Food, user models.User) (err error) {

	c := config.Config{}
	serverKey, _ := c.GetKey()

	var message string
	var title string

	if user.Lang == "pt" {
		message = fmt.Sprintf("O seu alimento %v vencerá na data %v. Esteja atento, abraços.", food.Description, food.LimitDate)
		title = "Olá, preciso te falar algo"
	} else {
		t, err := time.Parse("02/01/2006", food.LimitDate)
		if err != nil {
			fmt.Println(err.Error())
		}
		message = fmt.Sprintf("Your food %v will be outdated %v. Be aware and don't forget :)", food.Description, t.Format("01/02/2006"))
		title = "Hey, i need to tell you something"
	}

	msg := &fcm.Message{
		To: user.Token,
		Notification: &fcm.Notification{
			Body: message,
			Title: title,
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
