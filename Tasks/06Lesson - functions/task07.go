package main

import "fmt"

func main() {
	maxNum := maxOfThree(5, 11, 10)
	fmt.Println(maxNum)
}

func maxOfThree(a, b, c int) int {
	max := a

	if max < b {
		max = b
	}
	if max < c {
		max = c
	}
	return max
}
