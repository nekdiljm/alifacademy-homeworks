package main

import "fmt"

func main() {
	a := isAdult(18)
	fmt.Println(a)
}

func isAdult(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}
