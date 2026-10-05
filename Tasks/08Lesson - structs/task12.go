package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func main() {
	rectangle := Rectangle{
		Width:  23,
		Height: 11,
	}

	area := rectangle.Width * rectangle.Height
	perimeter := 2 * (rectangle.Width + rectangle.Height)

	fmt.Printf("Площать прямоугольника: %.1f\n", area)
	fmt.Printf("Периметр прямоугольника: %.1f\n", perimeter)
}
