package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Print(i)
		}(i)

	}
	wg.Wait()
}
