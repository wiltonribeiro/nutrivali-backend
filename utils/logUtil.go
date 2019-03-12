package utils

import (
	"go-app/models"
	"time"
)

var log *models.Log

func UpdateLog(content string, date string){
	log = &models.Log{Content: content, Value: date}
}

func GetLog() models.Log {
	if log == nil {
		return models.Log{Content: "waiting to start", Value: time.Now().String()}
	}
	return *log
}