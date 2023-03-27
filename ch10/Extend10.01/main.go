package main

import (
	"fmt"
	"time"
)

func deadline(t time.Time, tz string) (time.Time, error) {
	after5years := t.AddDate(+5, 0, 0)
	timezone, err := time.LoadLocation(tz)
	if err != nil {
		return t, err
	}
	deadlineTime := after5years.In(timezone)
	return deadlineTime, nil

}
func main() {
	testDate := time.Date(2021, 4, 23, 16, 18, 15, 1, time.Local)
	deadlineTime, err := deadline(testDate, "Asia/Tokyo")
	if err == nil {
		fmt.Println("curret time\t: " + testDate.Format("15:04:05 1/2/2006"))
		fmt.Println("deadline\t: " + deadlineTime.Format("15:04:05 1/2/2006"))
	} else {
		fmt.Println(err)
	}
}
