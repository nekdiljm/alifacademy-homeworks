package main

import "fmt"

type Address struct {
	City   string
	Street string
	Zip    string
}

type Company struct {
	Name    string
	Address Address
}

type Employee struct {
	Name    string
	Company Company
}

func main() {
	employee := Employee{
		Name: "Nekruz",
		Company: Company{
			Name: "RedCore",
			Address: Address{
				City:   "Dushanbe",
				Street: "Vozehh",
				Zip:    "88",
			},
		},
	}

	fmt.Printf("Город в котором работает сотрудник: %v\n", employee.Company.Address.City)
}
