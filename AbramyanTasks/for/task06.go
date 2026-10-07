package main

import "fmt"

/*
Дано вещественное число — цена 1 кг конфет. Вывести стоимость 1.2,
1.4, . . . , 2 кг конфет.
*/
func main() {
	price := 100.0

	for i := 11; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		a := (float64(i) * price) / 10
		fmt.Println(a)
	}
}
