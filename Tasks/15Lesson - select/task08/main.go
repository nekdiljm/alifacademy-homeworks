package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				even <- i
			}

			if i%2 != 0 {
				odd <- i
			}
		}

		close(even)
		close(odd)
	}()

	var evenClosed bool
	var oddClosed bool
outer:
	for {
		select {
		case v, ok := <-even:
			if !ok {
				evenClosed = true
				even = nil
				break
			}
			fmt.Println("четное:", v)
		case v, ok := <-odd:
			if !ok {
				oddClosed = true
				odd = nil
				break
			}
			fmt.Println("нечетное:", v)
		}

		if evenClosed && oddClosed {
			break outer
		}
	}

}
