package main

import "fmt"

/*
Дано вещественное число — цена 1 кг конфет. Вывести стоимость 0.1,
0.2, . . . , 1 кг конфет.
*/
func main() {
	var price int

	for i := 1; i <= 10; i++ {
		a := (float64(price) * float64(i)) / 10
		fmt.Println(a)
	}

}
