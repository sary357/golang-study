package main

import "fmt"

type Developer struct {
	Name       string
	HourlyRate uint
	WorkWeek   [7]uint8
}

type Weekday uint8

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) logHours(day Weekday, hours uint8) {
	d.WorkWeek[day] = hours
}

func (d Developer) WeekPayment() uint {
	var r uint = 0
	for _, v := range d.WorkWeek {
		r = d.HourlyRate*uint(v) + r
	}
	return r
}

func main() {
	developer := Developer{}
	developer.HourlyRate = 10
	developer.logHours(Sunday, 0)
	developer.logHours(Monday, 8)
	developer.logHours(Tuesday, 10)
	fmt.Printf("Working hours on Monday: %d.\n", developer.WorkWeek[Monday])
	fmt.Printf("Working hours on Tuesday: %d.\n", developer.WorkWeek[Tuesday])
	fmt.Printf("Weekly Wage: %d\n", developer.WeekPayment())
}
