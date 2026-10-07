package main

import "fmt"

func main() {
	var n int

	fmt.Println("Введите число")
	fmt.Scan(&n)

	found := false

	for i := 1; i <= n; i++ {
		if i%7 == 0 && i%3 == 0 {
			fmt.Println(i)
			found = true
		}

	}
	if !found {
		fmt.Println("Не найдено!")
	}
}
