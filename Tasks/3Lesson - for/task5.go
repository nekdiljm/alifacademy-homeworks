package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите число - ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%10 == 5 {
			continue
		}
		fmt.Println(i)
	}

}
