package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var sum int

	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}
