package main

import (
	"fmt"
)

func main() {
	var (
		fuelType     string
		numberLiters int
		price        int
	)

	fmt.Println("Введите тип топлива(92, 95, 98, diesel)")
	fmt.Scan(&fuelType)

	fmt.Print("Введите количество литров - ")
	fmt.Scan(&numberLiters)

	if fuelType == "92" {
		price = 10
	} else if fuelType == "95" {
		price = 12
	} else if fuelType == "98" {
		price = 15
	} else if fuelType == "diesel" {
		price = 11
	} else {
		fmt.Println("Неверный вид топлива!")
		return
	}
	fmt.Printf("Цена за литр - %d\n", price)
	fmt.Printf("Количество литров - %d\n", numberLiters)
	fmt.Printf("Итоговая стоимость - %d\n", price*numberLiters)
}
