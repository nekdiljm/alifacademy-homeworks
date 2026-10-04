package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Кол-во горутин до:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	for i := 1; i <= 50; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(100 * time.Millisecond)
		}()
	}
	fmt.Println("Кол-во горутин после цикла(до wg.Wait()):", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Кол-во горутин после цикла(после wg.Wait()):", runtime.NumGoroutine())
}
