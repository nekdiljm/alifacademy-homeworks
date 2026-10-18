package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	messages := make(chan string)

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		after := time.After(2 * time.Second)
		for {
			select {
			case <-ticker.C:
				messages <- "ping"
			case <-after:
				close(done)
				return
			}
		}
	}()

outer:
	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		case _, ok := <-done:
			if !ok {
				fmt.Println("Конец")
				break outer
			}
		}

	}
}
