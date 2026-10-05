package main

import (
	"fmt"

	"task1/geometry"
)

func main() {
	rectangleArea1 := geometry.RectangleArea(20, 10)
	rectanglePerimeter1 := geometry.RectanglePerimeter(20, 10)
	fmt.Printf("Площадь первого прямоугольника: %.1f\n", rectangleArea1)
	fmt.Printf("Периметр первого прямоугольника: %.1f\n\n", rectanglePerimeter1)

	rectangleArea2 := geometry.RectangleArea(10, 5)
	rectanglePerimeter2 := geometry.RectanglePerimeter(10, 5)
	fmt.Printf("Площадь второго прямоугольника: %.1f\n", rectangleArea2)
	fmt.Printf("Периметр второго прямоугольника: %.1f\n", rectanglePerimeter2)

}
