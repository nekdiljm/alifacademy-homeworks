package main

import "fmt"

func sendName(sl []string, ch chan string) {
	for _, v := range sl {
		ch <- v
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	students := []string{"Аня", "Борис", "Вика", "Данил", "Ева"}

	go sendName(students, ch)

	sl := make([]string, 0)
	var counter int

	for v := range ch {
		sl = append(sl, v)
		counter++
	}

	fmt.Printf("Имен получено: %d %v\n", counter, sl)
}
