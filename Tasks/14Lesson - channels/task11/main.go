package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i * i
		}
		close(ch)
	}()

	var total int
	for v := range ch {
		total += v
	}

	fmt.Println("Сумма всех чисел:", total)
}
