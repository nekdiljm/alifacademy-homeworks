package main

import (
	"fmt"
)

func main() {
	firstNum, secondNum := 2, 0
	divide := safeDivide(firstNum, secondNum)
	if divide != nil {
		fmt.Printf("%v / %v = %v", firstNum, secondNum, *divide)
	} else {
		fmt.Println("На ноль делить нельзя!")
	}

}

func safeDivide(a, b int) *int {
	if b == 0 {
		return nil
	}

	c := a / b
	return &c
}
