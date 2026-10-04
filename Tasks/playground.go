package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	//var ch chan int
	var ch = make(chan int)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go func() {
		defer close(ch)
		for _, v := range slice {
			select {
			case ch <- v:
			case <-ctx.Done():
				return
			}
		}
	}()

	const maxGo = 2
	wg := sync.WaitGroup{}

	for i := 1; i <= maxGo; i++ {
		wg.Go(func() {
			for {
				select {
				case v, ok := <-ch:
					if !ok {
						return
					}
					fmt.Printf("worker is done: %d, value is: %d\n", i, v)
				case <-ctx.Done():
					return
				}
			}
		})
	}

	wg.Wait()
}
