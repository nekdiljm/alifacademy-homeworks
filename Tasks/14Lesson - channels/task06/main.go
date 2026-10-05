package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	initTime := time.Now()

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Замер первой записи:", time.Since(initTime))

	ch2 := make(chan int, 5)
	initTime2 := time.Now()

	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	ch2 <- 4
	ch2 <- 5
	go func() {
		time.Sleep(200 * time.Millisecond)
		<-ch2
		time.Sleep(200 * time.Millisecond)
		<-ch2
		time.Sleep(200 * time.Millisecond)
		<-ch2
		time.Sleep(200 * time.Millisecond)
		<-ch2
		time.Sleep(200 * time.Millisecond)
		<-ch2
		time.Sleep(200 * time.Millisecond)
		<-ch2
	}()
	ch2 <- 6

	fmt.Println("Замер второй записи:", time.Since(initTime2))
}

// В буфер можно положить 5 значений без ожидания, но для 6 нужно, чтобы из канала сначала прочитали какое-нибудь значение
