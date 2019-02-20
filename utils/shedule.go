package utils

import (
	"github.com/jasonlvhit/gocron"
)

type task func()

func ScheduleByTime(time string,task task) {
	s1 := gocron.NewScheduler()

	s1.Every(1).Day().At(time).Do(func() {
		task()
	})

	<- s1.Start()
}
