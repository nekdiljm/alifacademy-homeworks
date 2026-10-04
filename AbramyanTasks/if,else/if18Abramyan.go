package main

import "fmt"

/*
Даны три целых числа, одно из которых отлично от двух других, равных между собой. Определить порядковый номер числа, отличного от
остальных
*/
func main() {
	var (
		a = 10
		b = 3
		c = 10
	)

	if a != b && a != c {
		fmt.Println("1")
	} else if b != a && b != c {
		fmt.Println("2")
	} else if c != a && c != b {
		fmt.Println("3")
	}
	/*2 вариант
	switch {
	case a != b && a != c:
		fmt.Println("1")
	case b != a && b != c:
		fmt.Println("2")
	case c != a && c != b:
		fmt.Println("3")
	}
	*/
}
