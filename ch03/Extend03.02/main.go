package main

import (
	"errors"
	"fmt"
)

func calculate_loan(credit_num int, monthly_salary int, loan_money int, loan_period int) (float32, int, error) {

	var annual_percentage int
	var monthly_paid float32
	var e error
	// verify input data and return error if anything wrong.
	if credit_num < 0 {
		e = errors.New("信用分數小於0")
	}
	if monthly_salary < 0 {
		e = errors.New("每月收入小於0")
	}
	if loan_money < 0.0 {
		e = errors.New("貸款金額小於0")
	}
	if loan_period < 0.0 {
		e = errors.New("貸款期數小於0")
	}
	if loan_period%12 != 0 {
		e = errors.New("貸款期數不能被12整除")
	}

	if credit_num >= 450 {
		annual_percentage = 15
		// calculate monthly loan
		monthly_paid = (float32(loan_money) + float32(loan_money)*(float32(annual_percentage)/100.0)) / float32(loan_period)
		if monthly_paid > float32(monthly_salary)/5.0 {
			e = errors.New("月付款金額超過收入20%")
		}
	} else {
		annual_percentage = 20
		// calculate monthly loan
		monthly_paid = (float32(loan_money) + float32(loan_money)*(float32(annual_percentage)/100.0)) / float32(loan_period)
		if monthly_paid > float32(monthly_salary/10.0) {
			e = errors.New("月付款金額超過收入10%")
		}
	}
	return float32(loan_money) * (float32(annual_percentage) / 100.0), annual_percentage, e
}

func main() {
	//1st person
	credit_num := 500
	income := 1000
	loan := 1000
	period := 24
	//(credit_num int, monthly_salary int, loan_money int, loan_period int
	total_interest, percentage, e := calculate_loan(credit_num, income, loan, period)
	fmt.Println("\n申請人 1\n----------")
	fmt.Println("信用分數:\t", credit_num)
	fmt.Println("收入:\t", income)
	fmt.Println("貸款金額:\t", loan)
	fmt.Println("貸款利率:\t", percentage)
	fmt.Println("貸款期數:\t", period)
	fmt.Println("總利息:\t", total_interest)
	fmt.Println("每月利息:\t", total_interest/float32(period))
	if e == nil {
		fmt.Println("申請通過:\t", true)
	} else {
		fmt.Println("申請通過:\t", false)

	}

	// 2nd person
	credit_num = 350
	income = 1000
	loan = 10000
	period = 12
	total_interest, percentage, e = calculate_loan(credit_num, income, loan, period)
	fmt.Println("\n申請人 2\n----------")
	fmt.Println("信用分數:\t", credit_num)
	fmt.Println("收入:\t", income)
	fmt.Println("貸款金額:\t", loan)
	fmt.Println("貸款利率:\t", percentage)
	fmt.Println("貸款期數:\t", period)
	fmt.Println("總利息:\t", total_interest)
	fmt.Println("每月利息:\t", total_interest/float32(period))
	if e == nil {
		fmt.Println("申請通過:\t", true)
	} else {
		fmt.Println("申請通過:\t", false)
	}
}
