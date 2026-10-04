package main

import "fmt"

func main() {
	arr := make([]string, 0)

	arr = append(arr, "Hello ")
	fmt.Printf("Значения - %v, Длина - %v, Объем - %v\n", arr, len(arr), cap(arr))

	arr = append(arr, "My ")
	fmt.Printf("Значения - %v, Длина - %v, Объем - %v\n", arr, len(arr), cap(arr))

	arr = append(arr, "Friend!")
	fmt.Printf("Значения - %v, Длина - %v, Объем - %v\n", arr, len(arr), cap(arr))
}
