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
		// 15: 24 小時制表示"時", 如果用 3 表示用 12 小時制
		// 04: 用 2 位數表示"分", 如果用 4 代表 1 位數表示"分"
		// 05: 用 2 位數表示"秒", 如果用 5 代表 1 位數表示"秒"
		// 01: 用 2 位數表示"月", 如果用 1 代表 1 位數表示"月"
		// 02: 用 2 位數表示"日", 如果用 2 代表 1 位數表示"日"
		// 2006:  4 位數表示"年"

		fmt.Println("curret time (in Taipei time)\t: " + testDate.Format("15:04:05 01/02/2006"))
		fmt.Println("deadline (in Tokyo time)\t: " + deadlineTime.Format("15:04:05 01/02/2006"))
	} else {
		fmt.Println(err)
	}
}
