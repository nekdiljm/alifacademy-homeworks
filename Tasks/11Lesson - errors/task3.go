package main

import (
	"fmt"
)

func Withdraw(balance, amount float64) (float64, error) {
	if amount < 0 {
		return 0, fmt.Errorf("отрицательная сумаа снятия")
	}

	if amount > balance {
		return 0, fmt.Errorf("сумма снятие больше чем баланс. Баланс: %.1f Снятие: %1.f", balance, amount)
	}

	return balance - amount, nil
}

func main() {
	withdraw, err := Withdraw(1000, 500)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	fmt.Println("После первой операции:", withdraw)
	fmt.Println("")

	withdraw, err = Withdraw(1000, -100)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	fmt.Println("После второй операции:", withdraw)
	fmt.Println("")

	withdraw, err = Withdraw(1000, 2000)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	fmt.Println("После 3 операции:", withdraw)
}
