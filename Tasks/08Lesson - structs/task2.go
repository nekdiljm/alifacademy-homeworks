package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

func main() {
	car := Car{}

	fmt.Println(car)
}
