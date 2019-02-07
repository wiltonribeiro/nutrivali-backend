package main

import (
	"fmt"
	"go-app/controllers"
	"go-app/server"
	"time"
)

func main() {
	go func() {
		for {
			controller := controllers.NotificationController{}
			err := controller.RequestNotify()
			if err!=nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("rodou já %v", time.Now().String())
			time.Sleep(time.Hour*22)
		}
	}()
	
	server.InitServer()
}

