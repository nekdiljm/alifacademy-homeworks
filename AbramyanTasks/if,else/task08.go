package main

import "fmt"

// Даны два числа. Вывести вначале большее, а затем меньшее из них.
func main() {
	var (
		a = 1
		b = 20
	)

	if a < b {
		fmt.Println(b, a)
	} else {
		fmt.Println(a, b)
	}
	fmt.Println()
}
