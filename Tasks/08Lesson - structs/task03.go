package main

import "fmt"

type Student struct {
	Name   string
	Grade  int
	Passed bool
}

func main() {
	student := Student{
		Name: "Nekruz",
	}

	fmt.Println(student)
}
