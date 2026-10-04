package main

import "fmt"

/*
Даны два целых числа A и B (A < B). Найти сумму всех целых чисел
от A до B включительно.
*/
func main() {
	var (
		a = 10
		b = 20
		c int
	)

	for i := a; i <= b; i++ {
		c += i

	}
	fmt.Println(c)

}
