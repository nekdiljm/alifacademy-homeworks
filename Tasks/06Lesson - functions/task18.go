package main

import "fmt"

func main() {
	sliceOfFruits := makeFruits()
	fmt.Println(sliceOfFruits)
}

func makeFruits() []string {
	sliceOfFruits := []string{"Банан", "Яблоко", "Груша", "Вишня"}
	return sliceOfFruits
}
