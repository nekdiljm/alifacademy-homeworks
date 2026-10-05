package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func main() {
	products := []Product{
		{Name: "Банан", Price: 20},
		{Name: "Огурец", Price: 7},
		{Name: "Помидор", Price: 10},
		{Name: "Капуска", Price: 15},
		{Name: "Арбуз", Price: 35},
	}
	var total float64

	for _, v := range products {
		total += v.Price
	}

	avgPrice := total / float64(len(products))

	fmt.Printf("Общая сумма продуктов: %v\n", total)
	fmt.Printf("Средняя цена: %v\n", avgPrice)

}
