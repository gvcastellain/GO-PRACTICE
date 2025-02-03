package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Dev struct {
	Salary int
}

type Ceo struct {
	Salary         int
	CompanyProfits int
}

func (d *Dev) CalculateSalary() int {
	return d.Salary
}

func (c *Ceo) CalculateSalary() int {
	return c.Salary + c.CompanyProfits
}

func CompanyExpense(s []SalaryCalculator) int {
	expense := 0

	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}

	return expense
}

func main() {
	d := &Dev{
		Salary: 4000,
	}

	c := &Ceo{
		Salary:         7000,
		CompanyProfits: 2000,
	}

	fmt.Printf("Total expense %d", CompanyExpense([]SalaryCalculator{c, d}))
}
