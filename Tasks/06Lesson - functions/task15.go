package main

import "fmt"

func main() {
	x, y := midpoint(12, 15, 16, 17)

	fmt.Printf("Координаты x середины: %.1f\n", x)
	fmt.Printf("Координаты y середины: %.1f", y)
}

func midpoint(x1, y1, x2, y2 float64) (float64, float64) {
	x := (x1 + x2) / 2
	y := (y1 + y2) / 2

	return x, y
}
