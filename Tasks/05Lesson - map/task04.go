package main

import "fmt"

func main() {
	m := map[string]int{
		"Behzod": 29,
		"Nekruz": 25,
		"Nekdil": 18,
	}
	fmt.Printf("Ключ Nekruz - %v\n", m["Nekruz"])

	m["Nekruz"]++

	fmt.Printf("Мапа - %v\n", m)
	fmt.Printf("Ключ Nekruz - %v\n", m["Nekruz"])
}
