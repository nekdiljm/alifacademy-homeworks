package main

import "fmt"

func main() {
	printPrices(map[string]float64{
		"Помидор": 40,
		"Зелень":  20,
		"Огурец":  15,
		"Перец":   5,
	})
}

func printPrices(prices map[string]float64) {
	fmt.Println(prices)
}
