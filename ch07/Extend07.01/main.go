package main

import "fmt"

// ////////////// 1 ///////////////
type Employee struct {
	Id        int
	FirstName string
	LastName  string
}
type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}
type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}
type Payer interface {
	Pay() (string, float64)
}

// ////////////// 2 ///////////////
func initData() (Developer, Manager) {
	// 程式設計師的考核資料
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

	// 程式設計師
	d := Developer{
		Individual:        Employee{Id: 1, FirstName: "Eric", LastName: "Davis"},
		HourlyRate:        35,
		HoursWorkedInYear: 2400,
		Review:            employeeReview,
	}
	// 主管
	m := Manager{
		Individual:     Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"},
		Salary:         150000,
		CommissionRate: .07,
	}
	return d, m
}

// ////////// 3 ///////////////
func (m Manager) calculate() float64 {
	return m.Salary + (m.Salary * m.CommissionRate)
}

// //////////// 4 /////////////
func (d Developer) calculate() float64 {
	return d.HourlyRate * d.HoursWorkedInYear
}

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
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

// ///////////// 7 /////////////
func (d Developer) Pay() (string, float64) {
	return d.Individual.FullName(), d.calculate()
}

func (m Manager) Pay() (string, float64) {
	return m.Individual.FullName(), m.calculate()
}

// ////////////// 8 //////////
func (d Developer) ReviewRating(reviewStandard map[string]int) float64 {
	var totalRating float64 = 0.0
	r := d.Review
	for _, v := range r {
		switch v.(type) {
		case float32, float64, int, int32, int64:
			//fmt.Println("------------")
			intValue := v.(int)
			//fmt.Println(float64(intValue))
			//fmt.Println(totalRating)
			totalRating = totalRating + float64(intValue)
		case string:
			//fmt.Println(v)
			temp := reviewStandard[v.(string)]
			totalRating = totalRating + float64(temp)
		default:
			totalRating += 0
		}
	}
	return totalRating
}

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
	fmt.Printf("Employee Name: %v, Annual Pay: %v\n", d.Individual.FullName(), d.calculate())
	fmt.Printf("Manage's Name: %v, Annual Pay: %v\n", m.Individual.FullName(), m.calculate())
}
