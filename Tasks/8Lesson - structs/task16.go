package main

import "fmt"

type Address struct {
	City   string
	Street string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	person := Person{
		Name: "Behruz",
		Age:  27,
		Address: Address{
			City:   "Dushanbe",
			Street: "Somoni",
		},
	}
	fmt.Printf("Структура до изменения: %v\n", person)

	person.Address.Street = "Rudaki"
	fmt.Printf("Структура после изменения: %v\n", person)
}
