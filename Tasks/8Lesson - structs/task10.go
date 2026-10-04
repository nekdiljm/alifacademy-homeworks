package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	product := Product{
		Name:     "Банан",
		Price:    24,
		Quantity: 13,
	}

	total := product.Price * float64(product.Quantity)
	fmt.Printf("Общая стоимость: %.1f\n", total)
}
