package main

import "fmt"

func main() {
	printSlice([]int{1, 2, 3, 4, 5})
}

func printSlice(nums []int) {
	fmt.Println("Срез чисел: ", nums)
}
