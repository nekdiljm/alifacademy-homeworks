package main

import "fmt"

/*
Даны два целых числа A и B (A < B). Найти произведение всех целых
чисел от A до B включительно.
*/
func main() {
	var (
		a = 5
		b = 10
		c = 1
	)

	for i := a; i <= b; i++ {
		c *= i
	}
	fmt.Println(c)

}
