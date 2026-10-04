package main

import "fmt"

/*
Даны два целых числа A и B (A < B). Вывести в порядке убывания все
целые числа, расположенные между A и B (не включая числа A и B), а
также количество N этих чисел
*/

func main() {
	var n int

	a := 10
	b := 50

	for i := 49; i < b; i-- {
		if i <= a {
			break
		}
		n++
		fmt.Println(i)
	}
	fmt.Printf("Кол-во чисел - %d", n)
}
