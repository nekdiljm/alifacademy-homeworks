package main

import "fmt"

func main() {
	totalPr := applyDiscount([]float64{100, 150, 200}, 5)
	fmt.Println(totalPr)
}

func applyDiscount(prices []float64, percent float64) []float64 {

	for i, _ := range prices {
		prices[i] -= (prices[i] * percent) / 100
	}
	return prices
}
