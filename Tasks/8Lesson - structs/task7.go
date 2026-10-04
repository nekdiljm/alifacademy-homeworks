package main

import "fmt"

type City struct {
	Name       string
	Population int
	IsCapital  bool
}

func main() {
	val1 := City{
		"Moscow",
		10000000,
		true,
	}

	val2 := City{
		Population: 10000000,
		Name:       "Moscow",
		IsCapital:  true,
	}

	fmt.Printf("%+v\n", val1)
	fmt.Printf("%+v\n", val2)
}
