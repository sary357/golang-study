package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	ErrorNot9Digit       = errors.New("Not a 9-digit string")
	ErrorNotANumber      = errors.New("Not a valid number")
	ErrorPrefixZero      = errors.New("Prefix include \"000\"")
	ErrorNoValid4thDigit = errors.New("The 4th digit is not 7 or 9")
)

// rule 1
func is9Digit(ssn string) error {
	s := strings.Replace(ssn, "-", "", -1)
	if len(s) != 9 {
		return fmt.Errorf("Not a 9-digit string: %v", ssn)
	}
	return nil
}

// rule 2
func isNumber(ssn string) error {
	s := strings.Replace(ssn, "-", "", -1)
	_, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("Not a valid number: %v", ssn)
	}
	return nil
}

// rule 3
func isPrefixExcludeThreeZero(ssn string) error {
	s := strings.Replace(ssn, "-", "", -1)
	if s[:3] == "000" {
		return fmt.Errorf("Prefix include \"000\": %v", ssn)
	}
	return nil
}

// rule 4
func isValid4thDigit(ssn string) error {
	s := strings.Replace(ssn, "-", "", -1)
	if s[0:1] == "9" && s[3:4] != "7" && s[3:4] != "9" {
		return fmt.Errorf("The 4th digit is \"7\" or \"9\" if the 1st digit is \"9\": %v", ssn)
	}
	return nil
}

func main() {
	fmt.Println("ok")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	validateSSN := []string{"123-45-6789", "012-8-678", "000-12-0962", "999-33-3333", "087-65-4321", "123-45-zzzz"}
	for k, i := range validateSSN {
		log.Println(fmt.Sprintf("Verifying data %v from %v: %v", k+1, len(validateSSN), i))

		isOK := true
		err := is9Digit(i)
		if err != nil {
			isOK = false
			log.Println(err)
		}

		err = isNumber(i)
		if err != nil {
			isOK = false
			log.Println(err)
		}

		err = isPrefixExcludeThreeZero(i)
		if err != nil {
			isOK = false
			log.Println(err)
		}
		err = isValid4thDigit(i)
		if err != nil {
			isOK = false
			log.Println(err)
		}

		if isOK {
			log.Println(fmt.Sprintf("ssn: %v is valid.", i))
		}
	}
}
