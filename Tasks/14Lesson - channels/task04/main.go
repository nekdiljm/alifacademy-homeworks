package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Printf("До чтения:\nlen ch: %d, cap ch: %d\n", len(ch), cap(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Printf("После чтения:\nlen ch: %d, cap ch: %d\n", len(ch), cap(ch))

}
