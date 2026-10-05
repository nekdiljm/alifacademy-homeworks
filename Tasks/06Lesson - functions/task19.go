package main

import "fmt"

func main() {
	sliceAdd := addToSlice([]int{1, 2, 3, 4}, 5)
	fmt.Println(sliceAdd)
}

func addToSlice(nums []int, value int) []int {
	nums = append(nums, value)
	return nums
}
