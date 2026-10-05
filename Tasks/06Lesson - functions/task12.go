package main

import "fmt"

func main() {
	area, perimeter := rectangleParams(20, 10)
	fmt.Printf("Площадь прямоугольника: %.1f\nПериметр прямоугольника: %.1f", area, perimeter)
}

func rectangleParams(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return area, perimeter
}
