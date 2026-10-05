package main

import (
	"fmt"
	"time"
)

func printSLowly(name string) {
	fmt.Println("Начали:", name)

	time.Sleep(time.Second)

	fmt.Println("Готово:", name)
}

func main() {
	initTime := time.Now()

	go printSLowly("Nekdil")
	go printSLowly("Nekruz")
	go printSLowly("Behzod")

	fmt.Println("Заняло времени:", time.Since(initTime))
}

// Печатается одно из 3 имен раз в несколько повторений,
// в основном главная main горутина заканчивает свое выполнение быстрее чем доходит до выполнения дочерних горутин
