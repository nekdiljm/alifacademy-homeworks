package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func main() {
	input := `{"name":"Ноутбук", "price":4500.5}`

	var item Item
	if err := json.Unmarshal([]byte(input), &item); err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(item)
}

// {Ноутбук 4500.5 0}
// При создании структуры, все поля получают default value для типа
// Поле которое не получило значение с json, остается с нулевым значением
