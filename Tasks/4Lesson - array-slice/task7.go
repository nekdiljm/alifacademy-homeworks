package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 6}

	//for i, _ := range arr {
	//	arr[i] = len(arr) - 1
	//}
	//fmt.Println(arr)

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%v ", arr[i])
	}

}
