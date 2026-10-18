package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	var (
		counter  int
		counter2 int
	)

	for i := 1; i <= 5; i++ {
		select {
		case ch <- i:
			counter++
		default:
			fmt.Printf("Буфер полон, пропустили число: %d\n", i)
			counter2++
		}
	}

	fmt.Printf("Попало в канал: %d\nПропустили: %d\n", counter, counter2)
}
