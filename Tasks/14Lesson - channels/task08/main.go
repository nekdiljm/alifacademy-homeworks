package main

import "fmt"

func generateNumbers(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go generateNumbers(5, ch)

	for v := range ch {
		fmt.Println(v)
	}
}
