package main

import "fmt"

type Passport struct {
	Number   string
	IssuedBy string
}

type Citizen struct {
	Name     string
	Passport Passport
}

func main() {
	citizen := Citizen{
		Name: "Ulugbek",
	}

	fmt.Printf("%+v", citizen)
}
