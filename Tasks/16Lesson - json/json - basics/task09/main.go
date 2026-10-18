package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	person := Person{
		Name: "Ulugbek",
		Age:  60,
		Address: Address{
			City:   "New York",
			Street: "Wall street",
		},
	}

	data, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	var p Person
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(p.Address.City)

}
