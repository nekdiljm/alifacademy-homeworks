package main

import "fmt"

// Даны три целых числа. Найти количество положительных и количество
// отрицательных чисел в исходном наборе.
func main() {
	var (
		a        int = 123
		b        int = 34
		c        int = -45
		countpos int
		countneg int
	)

	if a > 0 {
		countpos++
	} else if a < 0 {
		countneg++
	}

	if b > 0 {
		countpos++
	} else if b < 0 {
		countneg++
	}

	if c > 0 {
		countpos++
	} else if c < 0 {
		countneg++
	}
	fmt.Printf("Кол-во положительных - %d\nКол-во отрицательных - %d", countpos, countneg)
}
