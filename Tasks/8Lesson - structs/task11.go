package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "Nikita",
		Age:  17,
	}

	if person.Age < 18 {
		fmt.Println("Несовершеннолетний")
	} else {
		fmt.Println("Совершеннолетний")
	}
}
