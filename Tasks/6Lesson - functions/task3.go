package main

import "fmt"

func main() {
	showMenu()
}

func showMenu() {
	a := "1. Эспрессо!"
	b := "2. Латте!"
	c := "3. Американо!"
	d := "4. Капучино!"
	e := "5. Раф"

	fmt.Println(a, b, c, d, e)
}
