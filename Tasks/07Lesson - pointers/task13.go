package main

import "fmt"

func main() {
	var (
		current float64 = 25
		delta   float64 = 1
		minTemp float64 = 0
		maxTemp float64 = 100
	)

	adjTemp := adjustTemperature(&current, delta, minTemp, maxTemp)
	fmt.Println(adjTemp)
}

func adjustTemperature(current *float64, delta, min, max float64) bool {
	*current += delta

	if *current < min {
		*current = min
		return true
	}
	if *current > max {
		*current = max
		return true
	}

	return false
}
