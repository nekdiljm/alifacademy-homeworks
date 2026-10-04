package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(n int, wg *sync.WaitGroup) {
		defer wg.Done()

		for i := 1; i <= n; i++ {
			fmt.Println(i)
			time.Sleep(10 * time.Millisecond)
		}

	}(5, &wg)

	wg.Add(1)
	go func(n int, wg *sync.WaitGroup) {
		defer wg.Done()

		for i := 1; i <= n; i++ {
			fmt.Println(i)
			time.Sleep(10 * time.Millisecond)
		}

	}(3, &wg)

	wg.Add(1)
	go func(n int, wg *sync.WaitGroup) {
		defer wg.Done()

		for i := 1; i <= n; i++ {
			fmt.Println(i)
			time.Sleep(10 * time.Millisecond)
		}

	}(7, &wg)

	wg.Wait()
}
