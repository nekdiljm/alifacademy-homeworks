package main

import (
	"fmt"
	"math"
)

const pi = 3.14

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	diam   float64
	radius float64
}

type Rectangle struct {
	a float64
	b float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (c Circle) Area() float64 {
	return (pi * c.diam * c.diam) / 4
}

func (c Circle) Perimeter() float64 {
	return 2 * pi * c.radius
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.a + r.b)
}

func (t Triangle) Area() float64 {
	polPerim := (t.a + t.b + t.c) / 2
	return math.Sqrt(polPerim * (polPerim - t.a) * (polPerim - t.b) * (polPerim - t.c))
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func main() {
	shapes := []Shape{
		Circle{diam: 10, radius: 5},
		Rectangle{a: 8, b: 4},
		Triangle{a: 3, b: 4, c: 5},
	}

	var totalArea float64
	for _, v := range shapes {
		totalArea += v.Area()
	}

	fmt.Printf("Суммарная площадь: %.2f\n", totalArea)
}
