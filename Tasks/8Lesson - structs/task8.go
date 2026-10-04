package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

func main() {
	account := Account{
		Owner:   "Behzod",
		Balance: 1000.0,
	}

	fmt.Printf("До пополнения: %.1f\n", account.Balance)
	account.Balance += 250
	fmt.Printf("После пополнения: %.1f\n", account.Balance)

}
