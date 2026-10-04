package main

import (
	"fmt"
	"time"
)

func printSlowly(name string) {
	fmt.Println("Начали:", name)

	time.Sleep(time.Second)

	fmt.Println("Готово:", name)
}

func main() {
	initTime := time.Now()

	printSlowly("Nekdil")
	printSlowly("Nekruz")
	printSlowly("Behzod")

	fmt.Println("Заняло времени:", time.Since(initTime))
}
