package main

import "fmt"

type Wheel struct {
	Size  int
	Brand string
}

type Car struct {
	Model      string
	FrontWheel Wheel
	RearWheel  Wheel
}

func main() {
	car := Car{
		Model: "Jaguar",
		FrontWheel: Wheel{
			Size:  22,
			Brand: "Michelin",
		},
		RearWheel: Wheel{
			Size:  19,
			Brand: "Michelin",
		},
	}

	if car.FrontWheel.Size > car.RearWheel.Size {
		fmt.Printf("Переднее колесо больше!")
	} else {
		fmt.Printf("Задние колесе больше!")
	}
}
