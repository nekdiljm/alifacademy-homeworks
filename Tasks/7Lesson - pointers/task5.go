package main

import "fmt"

func main() {
	num := 40
	workWithCopy := double(num)
	fmt.Printf("func double: %v\n", workWithCopy)

	fmt.Printf("До работы функции с указателем: %v\n", num)
	doublePtr(&num)
	fmt.Printf("После работы функции с указателем: %v\n", num)

}

func double(n int) int {
	n *= 2
	return n
}

func doublePtr(n *int) {
	*n *= 2
	fmt.Printf("func doublePtr: %v\n", *n)
}
