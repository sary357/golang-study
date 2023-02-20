package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName       = errors.New("invalid last name")
	ErrInvalidRoutingNumbers = errors.New("invalid routing number")
)

func isValidLastName(name string) (bool, error) {
	if len(name) == 0 {
		return true, ErrInvalidLastName
	}
	return false, nil
}

func isValidRouteNumber(routingNumber int) (bool, error) {
	if (routingNumber) <= 0 {
		return true, ErrInvalidRoutingNumbers
	}
	return false, nil
}

func main() {

	lastName := ""
	routingNumbers := -1

	result, err := isValidLastName(lastName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = isValidRouteNumber(routingNumbers)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
