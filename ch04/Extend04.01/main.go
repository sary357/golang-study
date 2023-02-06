package main

import (
	"fmt"
)

func initializeArray(arr [10]int) [10]int {
	var startValue int = 10
	for i := 0; i < len(arr); i++ {
		arr[i] = startValue + i
	}
	return arr
}

func main() {
	var arr [10]int
	arr = initializeArray(arr)
	fmt.Println(arr)
}
