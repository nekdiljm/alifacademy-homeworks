package catalog

import "library/book"

func FindByAuthor(books []book.Book, author string) []book.Book {
	result := []book.Book{}
	for _, v := range books {
		if v.Author == author {
			result = append(result, v)
		}
	}

	return result
}
