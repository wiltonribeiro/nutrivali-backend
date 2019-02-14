package utils

import (
	"fmt"
	"time"
)

type task func()

var normalized = false

func Schedule(hours int, minutes int, task task) chan int {
	for{
		if normalized {
			task()
			time.Sleep(time.Hour*12)
		} else {
			t := time.Now().Add(time.Hour*24)
			day := fmt.Sprintf("%02d", t.Day())
			month := fmt.Sprintf("%02d", int(t.Month()))
			year := t.Year()
			h := fmt.Sprintf("%02d", hours)
			m := fmt.Sprintf("%02d", minutes)
			timeToTaskString := fmt.Sprintf("%v-%v-%v %v:%v UTC", year, month, day, h, m)
			timeToTask, _ := time.Parse("2006-01-02 15:04 MST", timeToTaskString)
			delta := time.Now().Sub(timeToTask)
			valueToWait := -1 *(delta.Minutes() - (3*time.Hour.Minutes()))
			duration := time.Minute*time.Duration(valueToWait)
			UpdateLog("waiting", duration.String())
			time.Sleep(duration)
			normalized = true
		}
	}
}
