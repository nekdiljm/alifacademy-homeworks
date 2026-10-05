package main

import "fmt"

type Counter struct {
	Value int
}

func main() {
	counter := Counter{
		Value: 0,
	}

	counter.Value++
	fmt.Printf("Значение Value после первого увеличения: %d\n", counter.Value)

	counter.Value++
	fmt.Printf("Значение Value после второго увеличения: %d\n", counter.Value)

	counter.Value++
	fmt.Printf("Значение Value после третьего увеличения: %d\n", counter.Value)

}
