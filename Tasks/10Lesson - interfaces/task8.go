package main

import "fmt"

type Point struct {
	X, Y int
}

func Describe(v any) {
	fmt.Printf("Значение: %v, Тип: %T\n", v, v)
}

func main() {
	Describe("Hello")
	Describe(14)
	Describe(15.)
	Describe(true)
	Describe([]int{1, 2, 3})
	Describe(map[string]int{
		"Ключ": 14,
	})
	Describe(Point{})
}
