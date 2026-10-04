package main

import "fmt"

func main() {
	var salary float64

	fmt.Print("Введите вашу запрлату: ")
	fmt.Scan(&salary)

	if salary <= 0 {
		fmt.Println("Вы ввели неверное число!")
		return
	} else if salary < 3000 {
		salary -= salary * 0.05
		fmt.Println("Ваш налог составляет - 5%")
	} else if salary == 3000 || salary <= 10000 {
		salary -= salary * 0.10
		fmt.Println("Ваш налог составляет - 10%")
	} else if salary > 10000 {
		salary -= salary * 0.15
		fmt.Println("Ваш налог составляет - 15%")
	}
	fmt.Println("Ваша зарплата после вычета налога составляет - ", salary)
}
