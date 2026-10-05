package catalog

type Product struct {
	Name string
	Price float64
}

func TotalPrice(products []Product) float64 {
	var totalPrice float64

	for _, v := range products {
		totalPrice += v.Price
	}

	return totalPrice
}