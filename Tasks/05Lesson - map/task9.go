package main

import "fmt"

func main() {
	arr := []string{"Олег", "Сергей", "Олег", "Николай", "Алекс", "Николай"}

	m := make(map[string]bool)

	for _, v := range arr {
		m[v] = true
	}
	fmt.Println("Кол-во уникальных посетителей:", len(m))

}
