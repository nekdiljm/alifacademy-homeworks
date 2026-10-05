package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Category string
}

func main() {
	products := []Product{
		{Name: "T-Shirt", Price: 70, Category: "одежда"},
		{Name: "Fridge", Price: 3000, Category: "электроника"},
		{Name: "Phone", Price: 5000, Category: "электроника"},
		{Name: "Apple", Price: 10},
		{Name: "Pineapple", Price: 40},
		{Name: "Shirt", Price: 120, Category: "одежда"},
	}

	productsAftDisc := make([]float64, 0, len(products))

	//МОЕ РЕШЕНИЕ
	for _, v := range products {

		if v.Category == "электроника" && v.Price > 1000 {
			v.Price -= (v.Price * 15) / 100
			productsAftDisc = append(productsAftDisc, v.Price)

		} else if v.Category == "одежда" {
			v.Price -= (v.Price * 10) / 100
			productsAftDisc = append(productsAftDisc, v.Price)
		} else {
			productsAftDisc = append(productsAftDisc, v.Price)
		}
	}

	// РЕШЕНИЕ ИИ
	//for _, v := range products {
	//	p := v
	//
	//	if v.Category == "электроника" && v.Price > 1000 {
	//		p.Price -= (v.Price * 15) / 100
	//	} else if v.Category == "одежда" {
	//		p.Price -= (v.Price * 10) / 100
	//	}
	//	productsAftDisc = append(productsAftDisc, p)
	//}
	//
	for _, v := range productsAftDisc {
		fmt.Printf("%+v\n", v)
	}
}
