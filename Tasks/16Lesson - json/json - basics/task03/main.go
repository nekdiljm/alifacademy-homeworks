package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	InStock bool    `json:"in_stock"`
}

func main() {
	product := Product{
		ID:      1,
		Name:    "Apple",
		Price:   55,
		InStock: true,
	}

	data, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(data))

}
