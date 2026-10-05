package main

import "fmt"

func main() {
	total := discountPrice(120, 5)
	fmt.Println(total)
}

func discountPrice(price, percent float64) float64 {
	discount := (price * percent) / 100
	return price - discount
}
