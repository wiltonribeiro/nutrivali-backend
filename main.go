package main

import (
	"go-app/controllers"
	"go-app/server"
	"go-app/utils"
	"time"
)

//para testar depois
func notificationSystem() {
	controller := controllers.NotificationController{}
	err := controller.RequestNotify()
	if err!=nil {
		utils.UpdateLog(err.Error(),time.Now().String())
	} else {
		utils.UpdateLog("running well", time.Now().String())
	}
}

func main() {
	go func() {
		for {
			controller := controllers.NotificationController{}
			err := controller.RequestNotify()
			if err!=nil {
				utils.UpdateLog(err.Error(),time.Now().String())
			} else {
				utils.UpdateLog("running well", time.Now().String())
			}

			time.Sleep(time.Hour*12)
		}
	}()
	server.InitServer()
}

