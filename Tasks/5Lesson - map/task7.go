package main

import "fmt"

func main() {
	m := map[string]int{
		"Абдурахим": 2,
		"Тимур":     5,
		"Самир":     3,
		"Некдил":    4,
	}

	for k, v := range m {
		fmt.Printf("Ученик - %s, Оценка - %d\n", k, v)
	}

	fmt.Println("")

	for k := range m {
		fmt.Printf("%s, ", k)
	}
}
