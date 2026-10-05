package main

import "fmt"

func main() {
	var (
		n int
		c int
	)

	fmt.Print("Введите число - ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			continue
		} else if i%7 == 0 {
			break
		}
		c += i

	}
	fmt.Println(c)

}
