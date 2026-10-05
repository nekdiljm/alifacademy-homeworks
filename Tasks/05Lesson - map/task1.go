package main

import "fmt"

func main() {
	m := map[string]int{
		"Вася":  4,
		"Петя":  5,
		"Гриша": 2,
	}

	fmt.Printf("Мапа - %v\n", m)
	fmt.Printf("Оценка Гриши - %v", m["Гриша"])
}
