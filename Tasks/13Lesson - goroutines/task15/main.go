package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	results := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			results[id] = 1
		}(i)
	}
	wg.Wait()

	var totalResult int
	for _, v := range results {
		totalResult += v
	}
	fmt.Println(totalResult)
}
