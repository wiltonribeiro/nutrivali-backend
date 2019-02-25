package utils

import (
	"go-app/models"
	"time"
)

var log *models.Log

func UpdateLog(content string, date string){
	log = &models.Log{Content: content, Date: date}
}

func GetLog() models.Log {
	if log == nil {
		return models.Log{Content: "waiting to start", Date: time.Now().String()}
	}
	return *log
}