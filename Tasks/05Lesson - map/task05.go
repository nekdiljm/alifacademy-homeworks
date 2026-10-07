package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Сыр"] = 4
	m["Колбаса"] = 5
	m["Кетчуп"] = 10
	m["Майонез"] = 0

	delete(m, "Майонез")

	_, ok := m["Майонез"]
	if !ok {
		fmt.Println("Товар больше нет")
	} else {
		fmt.Println("Товар найден Майонез")
	}
}
