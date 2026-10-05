package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	fmt.Println(arr)
}
