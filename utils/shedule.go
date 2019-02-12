package utils

type task func()

func schedule(hours int, minutes int, t task) chan int {
	var normalized = false
	for{
		if normalized{
			t()
			time.Sleep(time.Second*10)
		} else {
			t := time.Now().Add(time.Hour*24)
			day := t.Day()
			month := fmt.Sprintf("%02d", int(t.Month()))
			year := t.Year()
			timeToTaskString := fmt.Sprintf("%v-%v-%v %v:%v UTC", year, month, day, hours, minutes)
			timeToTask, _ := time.Parse("2006-01-02 15:04 MST", timeToTaskString)
			delta := time.Now().Sub(timeToTask)
			valueToWait := -1 *(delta.Minutes() - (3*time.Hour.Minutes()))
			duration := time.Minute*time.Duration(valueToWait)
			time.Sleep(duration)
			fmt.Printf("Sai do sleep as %v", time.Now().String())
			normalized = true
		}
	}
}
