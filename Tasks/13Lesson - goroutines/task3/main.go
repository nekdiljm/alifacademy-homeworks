package main

import (
	"fmt"
	"sync"
)

func printSLowly(wg *sync.WaitGroup, name string) {
	defer wg.Done()

	fmt.Println("Начали:", name)
	fmt.Println("Готово:", name)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printSLowly(&wg, "Nekdil")

	wg.Add(1)
	go printSLowly(&wg, "Nekruz")

	wg.Add(1)
	go printSLowly(&wg, "Behzod")

	wg.Wait()
}
