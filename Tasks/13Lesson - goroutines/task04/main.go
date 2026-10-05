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

	time.Sleep(2 * time.Second)
	fmt.Println("Заняло времени:", time.Since(initTime))
}

// Минус в том, что мы не знаем точное время завершения горутин, поэтому можем ждать лишнее время или наоборот
// завершить программу раньше, чем они закончат работу
