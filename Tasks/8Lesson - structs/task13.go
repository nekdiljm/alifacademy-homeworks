package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func main() {
	employee := Employee{
		Name:   "Nekruz",
		Salary: 50000.0,
	}
	fmt.Printf("Зарплата до повышения: %.1f\n", employee.Salary)

	salaryUp := (employee.Salary * 15) / 100
	employee.Salary += salaryUp
	fmt.Printf("Зарплата после повышения: %.1f\n", employee.Salary)

}
