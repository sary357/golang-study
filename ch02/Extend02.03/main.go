package main

import (
	"fmt"
)

func bubblesort(sample_slice []int) {
	is_changed := false
	var temp int
	for {
		is_changed = false
		for i := 0; i < len(sample_slice)-1; i++ {
			if sample_slice[i] > sample_slice[i+1] {
				temp = sample_slice[i+1]
				sample_slice[i+1] = sample_slice[i]
				sample_slice[i] = temp
				is_changed = true
			}
		}
		if is_changed == false {
			break
		}

	}
}

func main() {
	fmt.Println("ok")
	sample_slice := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println("排序前: ", sample_slice)
	bubblesort(sample_slice)
	fmt.Println("排序後: ", sample_slice)
}
