package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

func main() {
	book := Book{
		Title:  "War and Peace",
		Author: "Tolstoy",
		Year:   1869,
		Price:  100,
	}

	data, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Errors", err)
		return
	}

	fmt.Println(string(data))
}
