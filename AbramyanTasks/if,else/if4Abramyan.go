package main

import (
	"fmt"
)

// Даны три целых числа. Найти количество положительных чисел в исходном наборе
func main() {
	var (
		a     int = -2
		b     int = 6
		c     int = 7
		count int
	)

	if a > 0 {
		count++
	}
	if b > 0 {
		count++
	}
	if c > 0 {
		count++
	}
	fmt.Printf("Количество положительных чисел - %d", count)
}
