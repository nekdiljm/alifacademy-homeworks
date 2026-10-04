package main

import "fmt"

// Даны три числа. Найти наименьшее из них.
func main() {
	var (
		a = 15
		b = 7
		c = 9
	)

	if a < b || a < c {
		fmt.Println(a)
	} else if b < a || b < c {
		fmt.Println(b)
	} else if c < a || c < b {
		fmt.Println(c)
	}
	/* 2 вариант через switch
	switch {
	case a < b || a < c:
		fmt.Println(a)
	case b < a || b < c:
		fmt.Println(b)
	case c < a || c < b:
		fmt.Println(c)
	}
	*/
}
