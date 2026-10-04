package main

import "fmt"

// Дано целое число. Если оно является положительным, то прибавить к
// нему 1; в противном случае не изменять его. Вывести полученное число
func main() {
	var a int = 25

	if a > 0 {
		a++
		fmt.Println(a)
	} else {
		fmt.Println(a)
	}
}
