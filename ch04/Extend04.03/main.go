package main

import (
	"fmt"
)

func removeElement(strSlice []string, idx int) []string {
	var returnSlice []string
	if idx < 0 || idx > len(strSlice) {
		fmt.Println("Invalid idx:", idx)
		return strSlice
	}
	returnSlice = append(returnSlice, strSlice[0:idx]...)
	returnSlice = append(returnSlice, strSlice[idx+1:]...)
	return returnSlice
}

func main() {
	testSlice := []string{"Good", "Good", "Bad", "Good", "Good"}
	fmt.Println(removeElement(testSlice, 3))
}
