package main

import "fmt"

func main() {
	evenNums := filterEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(evenNums)
}

func filterEven(nums []int) []int {
	arr := make([]int, 0, len(nums))

	for _, v := range nums {
		if v%2 != 0 {
			continue
		}
		arr = append(arr, v)
	}
	return arr
}
