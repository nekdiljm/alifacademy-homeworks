package main

import "fmt"

func main() {
	num := new(int)
	*num = 100
	fmt.Printf("Результат: %v", *num)

}
