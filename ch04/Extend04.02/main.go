package main

import (
	"fmt"
)

func changeSeq(weekSlice []string) []string {
	var returnVal []string
	returnVal = append(returnVal, weekSlice[len(weekSlice)-1])
	returnVal = append(returnVal, weekSlice[:len(weekSlice)-1]...)
	return returnVal
}

func main() {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(changeSeq(week))
}
