package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Вывод до вызова функции: %v\n", arr)

	doubleArray(&arr)
	fmt.Printf("Вывод после вызова функции: %v", arr)
}

func doubleArray(arr *[5]int) {
	for i, _ := range arr {
		arr[i] *= 2
	}
}
