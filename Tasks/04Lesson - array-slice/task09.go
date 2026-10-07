package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var (
		even int
		odd  int
	)

	for _, v := range arr {
		if v%2 == 0 {
			even++
		} else if v%2 != 0 {
			odd++
		}
	}
	fmt.Printf("Кол-во четных - %v\nКол-во нечетных - %v", even, odd)
	//for _, v := range arr {
	//	switch {
	//	case v%2 == 0:
	//		even++
	//	default:
	//		odd++
	//	}
	//}
	//fmt.Printf("Кол-во четных - %v\nКол-во нечетных - %v", even, odd)

}
