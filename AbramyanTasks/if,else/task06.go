package main

import "fmt"

// Даны два числа. Вывести большее из них.
func main() {
	var (
		a int = 3
		b int = 2
	)

	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
