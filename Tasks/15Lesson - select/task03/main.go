package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)

		ch <- 42
	}()

	select {
	case msg := <-ch:
		fmt.Println("Дождались:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Таймаут")
	}
}
