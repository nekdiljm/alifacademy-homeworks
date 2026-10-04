package main

import "fmt"

type Book struct {
	Title string
	Pages int
	Price float64
}

func main() {
	book := []Book{
		{Title: "Праздник в Простоквашино", Pages: 120, Price: 55},
		{Title: "Romeо&Juliette", Pages: 500, Price: 250},
		{Title: "Собачье сердце", Pages: 30, Price: 30},
		{Title: "Война и Мир", Pages: 350, Price: 300},
	}

	maxPage := book[0].Pages
	for i, _ := range book {
		if maxPage < book[i].Pages {
			maxPage = book[i].Pages
		}
	}

	fmt.Printf("Максимальная страница: %v", maxPage)
}
