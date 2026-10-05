package main

import "fmt"

func main() {
	m := map[string]float64{
		"Яблоко":  12,
		"Колбаса": 20,
		"Гречка":  10,
		"Сыр":     50,
	}

	_, ok := m["Сыр"]
	if ok {
		fmt.Printf("Цена товара сыр = %v", m["Сыр"])
	} else {
		fmt.Println("Такого товара нет")
	}
}
