package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	fmt.Println("Введите 2 целых числа")
	fmt.Scan(&a, &b)
	//Первый вариант
	fmt.Printf("Сложение = %d\nВычитание = %d\nУмножение = %d\nДеление = %d\nОстаток от деления = %d\n ", a+b, a-b, a*b, a/b, a%b)

	/*
		Второй вариант
		//fmt.Println("Сложение =", a+b)
		fmt.Println("Вычитание =", a-b)
		fmt.Println("Умножение =", a*b)
		fmt.Println("Деление =", a/b)
		fmt.Println("Отсаток от деления =", a%b)
	*/
}
