package main

import (
	"fmt"
)

func calculator(tax_rate1 float64, tax_rate2 float64, tax_rate3 float64) float64 {
	return tax_rate1 + tax_rate2 + tax_rate3
}

func main() {
	cake_cost := 0.99
	cake_tax_rate := 0.075
	milk_cost := 2.75
	milk_tax_rate := 0.015
	cream_cost := 0.87
	cream_tax_rate := 0.02
	fmt.Println("Cake: ", cake_cost, fmt.Sprintf("%v", cake_tax_rate*100)+"%")
	fmt.Println("Milk: ", milk_cost, fmt.Sprintf("%v", milk_tax_rate*100)+"%")
	fmt.Println("Cream: ", cream_cost, fmt.Sprintf("%v", cream_tax_rate*100)+"%")
	fmt.Println("Sales Tax Total: ", calculator(cake_tax_rate, milk_tax_rate, cream_tax_rate))
}
