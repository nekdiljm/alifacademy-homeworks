package main

import "fmt"

func main() {
	age := 15
	ptr := &age

	fmt.Printf("Значение: %v\n", *ptr)
	fmt.Printf("Адрес в памяти: %v\n", &age)
	fmt.Printf("Тип адресы: %T\n", ptr)
}
