package main

import "fmt"

func main() {
	nameOfItem, maxPrice := mostExpensive(map[string]float64{
		"Apple":  20,
		"Oil":    50,
		"Flour":  55,
		"Tomato": 15,
	})
	fmt.Println(nameOfItem, maxPrice)
}

func mostExpensive(prices map[string]float64) (string, float64) {
	var (
		nameOfItem  string
		priceOfItem float64
	)

	for k, v := range prices {
		if priceOfItem < v {
			priceOfItem = v
			nameOfItem = k
		}
	}
	return nameOfItem, priceOfItem
}
