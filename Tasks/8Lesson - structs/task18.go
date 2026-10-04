package main

import "fmt"

type Dimensions struct {
	Width  float64
	Height float64
	Depth  float64
}

type Box struct {
	Name       string
	Dimensions Dimensions
}

func main() {
	box := Box{
		Name: "Box",
		Dimensions: Dimensions{
			Width:  25,
			Height: 15,
			Depth:  10,
		},
	}

	volume := box.Dimensions.Depth * box.Dimensions.Height * box.Dimensions.Width
	fmt.Printf("Объем коробки: %.1f", volume)
}
