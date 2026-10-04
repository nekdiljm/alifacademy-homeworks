package main

import "fmt"

/*
Дано вещественное число — цена 1 кг конфет. Вывести стоимость 1,
2, . . . , 10 кг конфет.
*/

func main() {
	price := 50.0

	for i := 1; i <= 10; i++ {
		a := price * float64(i)
		fmt.Println(a)
	}

}
