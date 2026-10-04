package main

import "fmt"

func main() {
	var weightKg float64
	var heightM float64

	fmt.Print("Введите ваш вес в кг: ")
	fmt.Scan(&weightKg)
	fmt.Print("Введите ваш рост в метрах: ")
	fmt.Scan(&heightM)

	imt := bmi(weightKg, heightM)
	fmt.Printf("Ваш индекс массы тела: %.1f", imt)
}

func bmi(weightKg, heightM float64) float64 {
	imt := weightKg / (heightM * heightM)
	return imt
}
