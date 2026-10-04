package main

import "fmt"

func main() {
	arr := []int{10, 50, 30, 40}
	var max int

	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
