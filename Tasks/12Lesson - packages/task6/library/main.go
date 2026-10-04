package main

import (
	"fmt"
	"library/book"
	"library/catalog"
)

func main() {
	books := []book.Book{
		{Title: "Crime and Punishment", Author: "Fyodor Dostoevsky", Year: 1866},
		{Title: "The Master and Margarita", Author: "Mikhail Bulgakov", Year: 1966},
		{Title: "The Little Prince", Author: "Antoine de Saint-Exupery", Year: 1943},
		{Title: "Crime and Punishment", Author: "Fyodor Dostoevsky", Year: 1866},
		{Title: "Heart of a Dog", Author: "Mikhail Bulgakov", Year: 1925},
	}

	findByAuthor := catalog.FindByAuthor(books, "Mikhail Bulgakov")

	for _, v := range findByAuthor {
		fmt.Printf("Title: %s || Author: %s || Year: %d\n", v.Title, v.Author, v.Year)
	}
}
