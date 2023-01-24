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

func getConfig() (bool, string, time.Time) { // exercise01.06
	return false, "info", time.Now()
}

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

	/////////////////////////////////
	// P.1-23
	// exercise01.06
	Debug1, LogLevel1, startUpTime1 := getConfig()
	fmt.Println(Debug1, LogLevel1, startUpTime1)
	/////////////////////////////////

	/////////////////////////////////
	// P.1-27
	// exercise01.07
	offset := 5
	fmt.Println(offset)

	offset = 10
	fmt.Println(offset)
	/////////////////////////////////

	/////////////////////////////////
	// P.1-29
	query, limit, offset := "bat", 10, 0
	query, maxlength, offset := "bat1", limit+10, 20
	fmt.Println(query, maxlength, offset)
	/////////////////////////////////

	var tip float64 = 20 - 1
	fmt.Println(tip)

	// P.1-32
	givenName := "Will"
	familyName := "Smith"
	fullName := givenName + " " + familyName
	fmt.Println(fullName)

	//P.1-33
	count := 5
	fmt.Println(count)
	count++
	fmt.Println(count)
	count += 5
	fmt.Println(count)

	//P.1-36
	visits := 15
	fmt.Println("New customer: ", visits == 1)
	fmt.Println("old customer: ", visits > 1)
	fmt.Println("silver customer: ", visits >= 11 && visits <= 20)
	fmt.Println("gold customer: ", visits >= 21 && visits <= 30)
	fmt.Println("premium customer: ", visits > 30)

	//P.1-38
	var count1 int
	fmt.Printf("Count: %#v \n", count1)
	var msg string
	fmt.Printf("msg: %#v \n", msg)
	var email []string
	fmt.Printf("email: %#v \n", email)
	var startTime time.Time
	fmt.Printf("time: %#v \n", startTime)
}
