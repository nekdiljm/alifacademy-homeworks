package main

import "fmt"

type Order struct {
	ID        int
	Amount    float64
	Completed bool
}

func main() {
	orders := []Order{
		{ID: 1, Amount: 150, Completed: true},
		{ID: 2, Amount: 60, Completed: false},
		{ID: 3, Amount: 30, Completed: true},
		{ID: 4, Amount: 50, Completed: true},
		{ID: 5, Amount: 100, Completed: false},
		{ID: 6, Amount: 10, Completed: true},
	}

	var (
		completedCount  int
		completedAmount float64
	)

	for _, v := range orders {
		if v.Completed {
			completedCount++
			completedAmount += v.Amount
		}
	}

	fmt.Printf("Завершенных заказов: %d\n", completedCount)
	fmt.Printf("Сумма заверешенных заказов: %.2f\n", completedAmount)

}
