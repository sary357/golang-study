package main

import (
	"fmt"

	"Extend08.01/payroll"
)

// ////////////// 1 ///////////////
// move to payroll

type Payer interface {
	Pay() (string, float64)
}

// ////////////// 2 ///////////////
func initData() (payroll.Developer, payroll.Manager) {
	// Developer's performance review data
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

	// Developer
	d := payroll.Developer{
		Individual:        payroll.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"},
		HourlyRate:        35,
		HoursWorkedInYear: 2400,
		Review:            employeeReview,
	}
	// Manager
	m := payroll.Manager{
		Individual:     payroll.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"},
		Salary:         150000,
		CommissionRate: .07,
	}
	return d, m
}

// ////////// 3 ///////////////

// //////////// 4 /////////////

// ///////////// 5 /////////////
func prepareReviewTable() map[string]int {
	return map[string]int{
		"Excellent":      5,
		"Good":           4,
		"Fair":           3,
		"Poor":           2,
		"Unsatisfactory": 1,
	}
}

// /////////// 6 //////////////

// ///////////// 7 /////////////

// ////////////// 8 //////////

// /////// 9 //////////
func payDetail(ps ...Payer) {
	for _, p := range ps {
		name, payAmount := p.Pay()
		fmt.Printf("%v Annual Pay this year: %v\n", name, payAmount)
	}
}
func main() {
	d, m := initData()
	reviewTable := prepareReviewTable()
	totoalReviewRating := d.ReviewRating(reviewTable)

	fmt.Printf("Employee Name: %v, Annual Rating: %v\n", d.Individual.FullName(), totoalReviewRating)
	fmt.Printf("Employee Name: %v, Annual Pay: %v\n", d.Individual.FullName(), d.Calculate())
	fmt.Printf("Manage's Name: %v, Annual Pay: %v\n", m.Individual.FullName(), m.Calculate())
}
