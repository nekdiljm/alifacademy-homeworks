package main

import "fmt"

func main() {
	m := make(map[string]float64, 4)
	m["Яблоко"] = 12
	m["Вишня"] = 20
	m["Апельсин"] = 30
	m["Арбуз"] = 40

	fmt.Printf("Мапа - %v\nДлина map - %v", m, len(m))
}
