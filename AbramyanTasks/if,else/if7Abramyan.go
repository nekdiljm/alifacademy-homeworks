package main

import "fmt"

// Даны два числа. Вывести порядковый номер меньшего из них.
func main() {
	var (
		a = 3
		b = 23
	)

	if a < b {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}
}
