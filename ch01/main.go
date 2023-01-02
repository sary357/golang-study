package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var foo string = "bar" //exercise01.02

var ( // exercise01.03
	Debug       bool      = false
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

var ( // exercise01.04
	Debug1       bool
	LogLevel1    = "info"
	startUpTime1 = time.Now()
)

func main() {
	/////////////////////////////////
	// exercise01.01
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
	/////////////////////////////////

	/////////////////////////////////
	//P.1-15
	//exercise01.02
	var baz string = "qux"
	fmt.Println(foo, baz)
	/////////////////////////////////

	/////////////////////////////////
	//P.1-16~1-17
	//exercise01.03
	fmt.Println(Debug, LogLevel, startUpTime)
	/////////////////////////////////

	/////////////////////////////////
	//P.1-18
	//exercise01.04
	fmt.Println(Debug1, LogLevel1, startUpTime1)
	/////////////////////////////////

	/////////////////////////////////
	//P.1-20
	//example01.03
	var seed int64 = 1234456789
	rand.Seed(seed)
	/////////////////////////////////

	/////////////////////////////////
	// P.1-21
	//exercise01.05
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
	/////////////////////////////////
}
