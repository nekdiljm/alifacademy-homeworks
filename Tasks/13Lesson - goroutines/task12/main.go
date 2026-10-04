package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	initTime := time.Now()
	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(1 * time.Millisecond)
		}()
	}
	wg.Wait()

	fmt.Println("Заняло времени:", time.Since(initTime))
}
