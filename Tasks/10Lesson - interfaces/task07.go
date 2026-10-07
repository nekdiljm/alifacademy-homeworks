package main

import "fmt"

type Discount interface {
	Apply(price float64) float64
}

type PercentDiscount struct {
	Percent float64
}

type FixedDiscount struct {
	Amount float64
}

type NoDiscount struct {
	Price float64
}

func (p PercentDiscount) Apply(price float64) float64 {
	return price - ((price * p.Percent) / 100)
}

func (f FixedDiscount) Apply(price float64) float64 {
	return price - f.Amount
}

func (n NoDiscount) Apply(price float64) float64 {
	return price
}

func main() {
	discounts := []Discount{
		PercentDiscount{Percent: 1},
		FixedDiscount{Amount: 5},
		NoDiscount{Price: 10},
	}

	price := 50.0
	for _, v := range discounts {
		price = v.Apply(price)
		fmt.Printf("Цена после: %v\n", price)

	}

}
