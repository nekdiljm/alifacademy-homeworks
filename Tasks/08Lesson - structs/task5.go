package main

import "fmt"

type Laptop struct {
	Brand string
	RAM   int
	SSD   int
	Price float64
}

func main() {
	laptop := Laptop{
		Brand: "MacBook",
		RAM:   16,
		SSD:   256,
		Price: 5000.0,
	}

	fmt.Printf("%s\n", laptop.Brand)
	fmt.Printf("%d\n", laptop.RAM)
	fmt.Printf("%d\n", laptop.SSD)
	fmt.Printf("%.1f\n", laptop.Price)
}
