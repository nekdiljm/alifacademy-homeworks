package main

import "fmt"

type Tempreture struct {
	Celsious float64
}

func main() {
	temp := Tempreture{
		Celsious: 25,
	}

	celsToFar := temp.Celsious*9/5 + 32

	fmt.Printf("Градусы Цельсия: %.1f\n", temp.Celsious)
	fmt.Printf("Градусы Фаренгейт: %.1f\n", celsToFar)
}
