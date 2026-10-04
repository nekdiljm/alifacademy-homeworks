package main

import "fmt"

type Employee interface {
	Salary() float64
	Position() string
}

type Manager struct {
	Name       string
	BaseSalary float64
}

func (m Manager) Salary() float64 {
	return m.BaseSalary + (m.BaseSalary*20)/100
}

func (m Manager) Position() string {
	return "Manager"
}

type Developer struct {
	Name       string
	BaseSalary float64
	Level      string
}

func (d Developer) Salary() float64 {
	if d.Level == "senior" {
		return d.BaseSalary + (d.BaseSalary*30)/100
	}

	return d.BaseSalary
}

func (d Developer) Position() string {
	return "Developer"
}

func main() {
	employeers := []Employee{
		Manager{Name: "Behruz", BaseSalary: 5000},
		Developer{Name: "Nekruz", BaseSalary: 20000, Level: "senior"},
		Manager{Name: "Anush", BaseSalary: 3500},
		Developer{Name: "Ashur", BaseSalary: 14000, Level: "middle"},
		Manager{Name: "Sadriddin", BaseSalary: 4500},
		Developer{Name: "Behzod", BaseSalary: 12000, Level: "junior"},
	}

	var bestSalary float64
	var bestEmployee Employee
	for _, v := range employeers {
		if bestSalary < v.Salary() {
			bestSalary = v.Salary()
			bestEmployee = v
		}
	}

	fmt.Printf("%v %.1f", bestEmployee.Position(), bestSalary)
}
