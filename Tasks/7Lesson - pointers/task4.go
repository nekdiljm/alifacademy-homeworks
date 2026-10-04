package main

import "fmt"

func main() {
	persEnergy := 100
	spendEnergy(&persEnergy, -1)
}

func spendEnergy(energy *int, cost int) {
	if cost > *energy {
		*energy = 0
	} else if cost < 0 {
		fmt.Println("Неверное число")
		return
	} else {
		*energy -= cost
	}

	fmt.Printf("Энергии осталось: %v", *energy)
}
