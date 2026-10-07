package main

import "fmt"

/*
Даны четыре целых числа, одно из которых отлично от трех других,
равных между собой. Определить порядковый номер числа, отличного от
остальных
*/
func main() {
	var (
		a = 5
		b = 5
		c = 5
		d = 8
	)

	if a != b && a != c && a != d {
		fmt.Println("1")
	} else if b != a && b != c && b != d {
		fmt.Println("2")
	} else if c != a && c != b && c != d {
		fmt.Println("3")
	} else if d != a && d != b && d != c {
		fmt.Println("4")
	}
	//2 вариант через switch
	//switch {
	//case a != b && a != c && a != d:
	//	fmt.Println("1")
	//case b != a && b != c && b != d:
	//	fmt.Println("2")
	//case c != a && c != b && c != d:
	//	fmt.Println("3")
	//case d != a && d != b && d != c:
	//	fmt.Println("4")
	//}

}
