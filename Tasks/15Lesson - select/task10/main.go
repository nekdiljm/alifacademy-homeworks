package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	results := make(chan int)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			time.Sleep(time.Second)
			results <- id
		}(i)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	var total int
	for v := range results {
		total += v
	}

	fmt.Println(total)
}
