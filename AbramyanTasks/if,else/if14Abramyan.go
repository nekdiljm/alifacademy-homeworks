package main

import "fmt"

// Даны три числа. Вывести вначале наименьшее, а затем наибольшее из
// данных чисел.
func main() {
	var (
		a int = 10
		b int = 5
		c int = 1
	)

	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	fmt.Println(min, max)

}
