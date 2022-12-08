package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var foo string = "bar" //Exercise 02

// Exercise 03
var (
	Debug       bool      = false
	Loglevel    string    = "info"
	startUptime time.Time = time.Now()
)

const (
	January = iota
	Feberaury
	March
)

// Exercise 06
func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func add5Point(count *int) {
	*count += 5
}
func add5Value(count int) {
	count += 5
}
func main() {

	// Exercise 01
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)

	//Exercise 02
	var baz string = "qux"
	fmt.Println(foo, baz)

	// Exercise 03
	fmt.Println(Debug, Loglevel, startUptime)

	// Exercise 04
	var seed int64 = 12344556789 //added int64
	rand.Seed(seed)

	// Exercise 05
	Debug := false
	Loglevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, Loglevel, startUpTime)

	// Example 01.04
	Debug, Loglevel, startUpTime = false, "info", time.Now()
	fmt.Println(Debug, Loglevel, startUpTime)

	// Exercise 06
	Debug, Loglevel, startUpTime = getConfig()
	fmt.Println(Debug, Loglevel, startUpTime)

	var offset int = 5
	fmt.Println(offset)
	defaultOffset := 10
	offset1 := 10
	offset1 = offset1 + defaultOffset
	fmt.Println(offset1)

	var count1 *int
	count2 := new(int)
	*count2 = 10
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	if count1 != nil {
		fmt.Printf("count1: %#v\n", count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", count2)
		fmt.Printf("*count2: %#v\n", *count2)
	}
	if count3 != nil {

		fmt.Printf("count3: %#v\n", count3)
		fmt.Printf("*count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("t: %#v\n", t)
		fmt.Printf("time: %#v\n", *t)
		fmt.Printf("time: %#v\n", t.String())
	}

	var count int = 10
	add5Point(&count)
	fmt.Println(count)
	add5Value(count)
	fmt.Println(count)

	fmt.Println(March)
}
