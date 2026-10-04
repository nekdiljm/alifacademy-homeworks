package main

import "fmt"

func main() {
	attemps := []float64{80, 50, 100, 74}
	best := attemps[0]
	for i, v := range attemps {
		attemptNum := i + 2
		if TrackRecord(&best, v) {
			fmt.Printf("Номер попытки - %v, Значение - %v\n", attemptNum, v)
		}
	}
	fmt.Printf("Лучшая попытка: %v\n", best)
}

func TrackRecord(best *float64, attemp float64) bool {
	if attemp > *best {
		*best = attemp
		return true
	}
	return false
}
