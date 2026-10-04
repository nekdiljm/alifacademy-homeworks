package main

import "fmt"

func main() {
	airTemp := 25
	ptr := &airTemp
	fmt.Printf("Температура до: %d\n", airTemp)

	*ptr = airTemp - 5
	fmt.Printf("Температура после: %v\n", *ptr)
}
