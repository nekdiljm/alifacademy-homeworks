package main

import "fmt"

// Даны три числа. Найти сумму двух наибольших из них.
func main() {
	var (
		a = 30
		b = 26
		c = 25
	)
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	fmt.Println(a + b + c - min)
}

//if a < b && a > c {
//	fmt.Println(a + b)
//} else if b < a && b > c {
//	fmt.Println(b + a)
//} else if c < a && c > b {
//	fmt.Println(c + a)
