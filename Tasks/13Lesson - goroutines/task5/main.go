package main

import (
	"fmt"
	"sync"
	"time"
)

func greet(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Привет, %s!\n", name)
	time.Sleep(300 * time.Millisecond)
}

func main() {
	names := []string{"Ali", "Umed", "Ulugbek", "Olim", "Aminjon"}
	var wg sync.WaitGroup

	for _, v := range names {
		wg.Add(1)
		go greet(v, &wg)
	}

	wg.Wait()
}
