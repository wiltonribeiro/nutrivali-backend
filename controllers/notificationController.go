package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"go-app/services"
	"go-app/utils"
	"time"
)

type NotificationController struct {}

func (u *NotificationController) Notify(ctx iris.Context) {

	var header = ctx.GetHeader("X-Appengine-Cron")

	if header == "" {
		ctx.StatusCode(401)
	} else {
		service := services.NotificationService{}
		count , err := service.RequestNotify()
		if err!=nil {
			utils.UpdateLog(err.Error(),time.Now().String())
		} else {
			utils.UpdateLog("notifications sent", fmt.Sprint(count))
		}

		ctx.StatusCode(200)
	}

}
