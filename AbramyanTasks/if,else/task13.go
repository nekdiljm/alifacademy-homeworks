package main

import "fmt"

// Даны три числа. Найти среднее из них (то есть число, расположенное
// между наименьшим и наибольшим).
func main() {
	var (
		a = 35
		b = 30
		c = 29
	)

	//if a > b || a < b && a < c || a > c {
	//	fmt.Println(a)
	//} else if b > a || b < a && b < c || b > c {
	//	fmt.Println(b)
	//} else if c > a || c < a && c < b || c > b {
	//	fmt.Println(c)
	//}
	//Вариант через switch
	switch {
	case (a < b && a > c) || (a > b && a < c):
		fmt.Println(a)
	case (b < a && b > c) || (b > a && b < c):
		fmt.Println(b)
	case (c < a && c > b) || (c > a && c < b):
		fmt.Println(c)
	}
}
