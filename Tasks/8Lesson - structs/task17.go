package main

import "fmt"

type Engine struct {
	HorsePower int
	Fuel       string
}

type Car struct {
	Brand  string
	Engine Engine
}

func main() {
	car := Car{
		Brand: "BMW",
		Engine: Engine{
			HorsePower: 350,
			Fuel:       "Gasoline",
		},
	}
	fmt.Printf("Мощность двигателя: %v лошадинных сил", car.Engine.HorsePower)
}
