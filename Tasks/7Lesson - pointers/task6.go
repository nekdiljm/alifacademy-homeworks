package main

import "fmt"

func main() {
	salary := 5000.
	bonus := 1000.
	applyBonus(&salary, bonus)

	fmt.Println("Общая зарплата за месяц =", salary)
}

func applyBonus(salary *float64, bonus float64) {
	*salary += bonus
	fmt.Printf("Зарплата с премией: %v\n", *salary)
}
