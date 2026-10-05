package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int)
	num := 7
	go square(num, ch)

	fmt.Printf("%d в квадрате = %d\n", num, <-ch)
}
