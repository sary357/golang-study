package main

import "fmt"

type closureFun func(int) int

type Developer struct {
	Name       string
	HourlyRate int
	WorkWeek   [7]int
	f          closureFun
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func nonLoggedHours() func(int) int {
	counter := 0
	return func(workHours int) int {
		if workHours == -1 {
			r := counter
			counter = 0
			return r
		} else {
			counter += workHours
			return counter
		}
	}
}

func (d *Developer) logHours() int {
	return d.f(-1)
}

func main() {
	developer := Developer{}
	developer.HourlyRate = 10
	developer.f = nonLoggedHours()
	developer.WorkWeek[Monday] = 8
	developer.f(8)
	developer.WorkWeek[Tuesday] = 7
	developer.f(7)
	developer.WorkWeek[Wednesday] = 10
	developer.f(10)
	totalWorkHours := developer.logHours()
	totalWage := developer.HourlyRate * totalWorkHours
	fmt.Printf("Working hours on Monday: %d.\n", developer.WorkWeek[Monday])
	fmt.Printf("Working hours on Tuesday: %d.\n", developer.WorkWeek[Tuesday])
	fmt.Printf("Working hours on Wednesday: %d.\n", developer.WorkWeek[Wednesday])
	fmt.Printf("Weekly Wage: %d\n", totalWage)
}
