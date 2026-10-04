package main

import "fmt"

// Даны целые числа K и N (N > 0). Вывести N раз число K
func main() {
	var (
		k = 7
		n = 5
	)

	for i := 0; i < n; i++ {
		fmt.Println(k)
	}
}
