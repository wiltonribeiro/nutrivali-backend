package main

import (
	"go-app/config"
	"go-app/server"
	"go-app/services"
	"go-app/utils"
	"time"
)

func notify(){
	service := services.NotificationService{}
	err := service.RequestNotify()
	if err!=nil {
		utils.UpdateLog(err.Error(),time.Now().String())
	} else {
		utils.UpdateLog("waiting next request", time.Now().String())
	}
}

func main() {
	go utils.ScheduleByTime("10:00", notify)
	go utils.ScheduleByTime("20:00", notify)

	c := config.Config{}

	if err := c.InitDB(); err == nil || config.DB != nil {
		server.InitServer()
	}

	defer c.CloseDB()

}

