package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("сумма пополнения отрицательная")
	}

	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("отрицательное число")
	}
	if amount > w.balance {
		return fmt.Errorf("сумма снятия %v ваш баланс: %v", amount, w.Balance())
	}

	w.balance -= amount
	return nil
}

func (w Wallet) Balance() float64 {
	return w.balance
}

func main() {
	wallet := Wallet{
		balance: 500,
	}

	fmt.Printf("Ваш баланс: %v\n", wallet.Balance())

	if err := wallet.Deposit(100); err != nil {
		fmt.Println("ошибка", err)
		return
	}
	fmt.Printf("Ваш баланс: %v\n", wallet.Balance())

	if err := wallet.Withdraw(300); err != nil {
		fmt.Println("ошибка", err)
		return
	}
	fmt.Printf("Ваш баланс: %v\n", wallet.Balance())

	if err := wallet.Deposit(-1); err != nil {
		fmt.Println("ошибка", err)
		return
	}
	fmt.Printf("Ваш баланс: %v\n", wallet.Balance())

}
