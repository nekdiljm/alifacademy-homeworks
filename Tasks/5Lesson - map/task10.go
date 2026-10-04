package main

import "fmt"

func main() {
	m := map[string][]string{
		"A1": {"Сергей", "Андрей", "Олег"},
		"A2": {"Максим", "Улугбек", "Иван"},
	}

	fmt.Printf("Мапа до добавления нового студента - %v, %v\n", m["A1"], m["A2"])

	m["A1"] = append(m["A1"], "Анюта")

	fmt.Println("")
	fmt.Printf("Мапа после добавления нового студента - %v, %v\n", m["A1"], m["A2"])
}
