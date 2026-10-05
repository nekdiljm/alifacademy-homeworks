package main

import "fmt"

type Engine struct {
	Power int
	Type  string
}

type Car struct {
	Brand  string
	Engine Engine
}

func main() {
	car1 := Car{
		Brand: "Mercedes",
		Engine: Engine{
			Power: 330,
			Type:  "Gasoline",
		},
	}

	car2 := Car{
		Brand: "BMW",
		Engine: Engine{
			Power: 350,
			Type:  "Diesel",
		},
	}

	if car1.Engine.Power > car2.Engine.Power {
		fmt.Printf("%s мощнее: %v лошадинных сил", car1.Brand, car1.Engine.Power)
	} else {
		fmt.Printf("%s мощнее: %v лошадинных сил", car2.Brand, car2.Engine.Power)
	}
}
