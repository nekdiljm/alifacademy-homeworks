package main

import "fmt"

func main() {
	var (
		n     int
		max   int
		digit int
	)

	fmt.Print("Введите число - ")
	fmt.Scan(&n)

	for n > 0 {
		digit = n % 10
		n /= 10
		if max < digit {
			max = digit
		}

	}
	fmt.Println(max)

}
