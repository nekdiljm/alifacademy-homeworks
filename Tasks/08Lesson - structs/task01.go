package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

func main() {
	book := Book{
		Title:  "Праздник в простоквашино",
		Author: "Эдуард Успенский",
		Year:   2004,
		Price:  53.5,
	}

	fmt.Printf("Название книги: %s\n", book.Title)
	fmt.Printf("Автор книги: %s\n", book.Author)
	fmt.Printf("Год выпуска: %d\n", book.Year)
	fmt.Printf("Цена: %.1f\n", book.Price)

	fmt.Printf("%+v", book)
}
