package main

import "fmt"

func main() {
	arr := make([]int, 2)

	for i := 0; i <= 14; i++ {
		arr = append(arr, i)
		fmt.Printf("len - %v, cap - %v\n", len(arr), cap(arr))
	}
}
