package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	initTime := time.Now()
	for i := 1; i <= 8; i++ {
		fmt.Printf("Цикл %d завершен\n", i)
		time.Sleep(time.Duration(i) * 50 * time.Millisecond)
	}
	fmt.Println("Обычный цикл занял:", time.Since(initTime))

	initTime2 := time.Now()
	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("Горутина %d заверешена\n", id)
			time.Sleep(time.Duration(id) * 50 * time.Millisecond)
		}(i)
	}
	wg.Wait()

	fmt.Println("Цикл через горутины занял:", time.Since(initTime2))

}
