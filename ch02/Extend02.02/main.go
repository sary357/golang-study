package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4}
	var max int = -1
	var max_word string = ""
	for k, v := range words {
		if v >= max {
			max = v
			max_word = k
		}
	}

	fmt.Println("出現最多次的字: ", max_word)
	fmt.Println("出現次數:      ", max)

}
