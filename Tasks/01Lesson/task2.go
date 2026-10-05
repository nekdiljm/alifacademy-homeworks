package main

import "fmt"

func main() {
	var (
		a int = 20
		b int = 5
	)
	sum := a + b
	sub := a - b
	multip := a * b
	divis := a / b
	modDivis := a % b
	fmt.Println(sum, sub, multip, divis, modDivis)
}
