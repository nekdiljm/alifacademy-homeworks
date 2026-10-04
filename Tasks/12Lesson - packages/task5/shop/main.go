package main

import (
	"fmt"
	"shop/catalog"
)

func main() {
	products := []catalog.Product{
		{Name: "Apple", Price: 20},
		{Name: "Pineapple", Price: 10},
		{Name: "Cheeze", Price: 30},
		{Name: "Garlic", Price: 50},
	}

	totalPrice := catalog.TotalPrice(products)
	fmt.Printf("Общая стоимость товаров: %.1f\n", totalPrice)
}
