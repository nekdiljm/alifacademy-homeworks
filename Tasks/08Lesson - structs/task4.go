package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	a := Point{
		X: 10,
		Y: 20,
	}

	b := Point{
		20,
		30,
	}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%v\n", b)
}
