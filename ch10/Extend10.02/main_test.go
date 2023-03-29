package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	loop := 1000
	for i := 0; i < loop; i++ {
		start := time.Now()
		value := Fibonacci(30)
		if value != 832040 {
			t.Fatal("Expected result: 832040. but the value we get is: " + strconv.Itoa(value))
		}
		end := time.Now()
		duration := end.Sub(start)
		fmt.Println(time.Now().Format("2006/01/02 15:04:05")+" ", duration)
	}
}
