package main

import (
	"fmt"
)

type Transaction struct {
	Type   string
	Amount float64
}

func main() {
	transactions := []Transaction{
		{Type: "deposit", Amount: 700},
		{Type: "deposit", Amount: 200},
		{Type: "withdraw", Amount: 1100},
		{Type: "withdraw", Amount: 100},
		{Type: "deposit", Amount: 100},
	}

	var (
		balance float64
		count   int
	)

	for _, v := range transactions {

		if v.Type == "deposit" {
			balance += v.Amount
		} else if v.Type == "withdraw" {
			if balance-v.Amount < 0 {
				count++
				continue
			}
			balance -= v.Amount
		}

		//switch v.Type {
		//case "deposit":
		//	balance += v.Amount
		//case "withdraw":
		//	if balance-v.Amount < 0 {
		//		count++
		//		continue
		//	}
		//	balance -= v.Amount
		//}

	}
	fmt.Printf("Ваш баланс: %.1f\nКол-во отклоненных операций: %v\n", balance, count)
}
