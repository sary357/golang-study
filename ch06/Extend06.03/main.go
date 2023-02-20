package main

import (
	"errors"
	"fmt"
)

type validateLastName func(string) error

type validateRoutingNumber func(int) error

type directDeposity struct {
	lastName              string
	firstName             string
	bankName              string
	rountingNumber        int
	accountNumber         int
	bankCode              int
	validateLastName      validateLastName
	validateRoutingNumber validateRoutingNumber
}

func isValidLastName(name string) error {
	if len(name) == 0 {
		return ErrInvalidLastName
	}
	return nil
}

func isValidRouteNumber(routingNumber int) error {
	if (routingNumber) <= 100 {
		panic(ErrInvalidRoutingNumbers)
	}
	return nil
}

var (
	ErrInvalidLastName       = errors.New("invalid last name")
	ErrInvalidRoutingNumbers = errors.New("invalid routing number")
)

func main() {
	someone := directDeposity{
		lastName:              "",
		firstName:             "Abe",
		bankName:              "XYZ inc.",
		accountNumber:         1809,
		rountingNumber:        17,
		bankCode:              17,
		validateLastName:      isValidLastName,
		validateRoutingNumber: isValidRouteNumber,
	}

	fmt.Println(someone.validateRoutingNumber(someone.rountingNumber))
	fmt.Println(someone.validateLastName(someone.lastName))
	fmt.Println("---------------")
	fmt.Printf("Last Name: %v\n", someone.lastName)
	fmt.Printf("First Name: %v\n", someone.firstName)
	fmt.Printf("Bank name: %v\n", someone.bankName)
	fmt.Printf("Bank code: %v\n", someone.bankCode)
	fmt.Printf("Account number: %v\n", someone.accountNumber)

}
