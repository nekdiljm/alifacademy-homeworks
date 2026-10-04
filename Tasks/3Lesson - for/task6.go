package main

import "fmt"

func main() {
	var (
		a   uint64
		sum uint64
	)

	fmt.Print("Введите положительное целое число - ")
	fmt.Scan(&a)

	for a > 0 {
		sum += a % 10
		a /= 10
	}
	fmt.Println(sum)

}
