package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main() {
	input := `{"title":"Go для начинающих", "author":"Иван Петров", "year": 2023}`

	var book Book
	if err := json.Unmarshal([]byte(input), &book); err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(book.Title)
	fmt.Println(book.Author)
	fmt.Println(book.Year)
}
