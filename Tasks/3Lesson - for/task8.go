package main

import "fmt"

func main() {
	var (
		a     int
		b     int
		digit int
	)

	fmt.Print("Введите число - ")
	fmt.Scan(&a)

	for a > 0 {
		digit = a % 10
		b = (b * 10) + digit
		a /= 10
	}
	fmt.Print(b)

}
