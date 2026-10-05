package main

import "fmt"

func main() {
	scoreOfPupils := makeScores()
	fmt.Println(scoreOfPupils)
}

func makeScores() map[string]int {
	m := map[string]int{
		"Аминчон": 5,
		"Собир":   3,
		"Умед":    4,
	}
	return m
}
