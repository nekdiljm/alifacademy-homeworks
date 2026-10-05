package main

import "fmt"

func main() {
	var (
		current   int
		totalLaps = 50
		finished  bool
	)

	for current = 0; current < totalLaps; {
		advanceLap(&current, totalLaps, &finished)
		fmt.Println(current, totalLaps, finished)

	}
}

func advanceLap(current *int, totalLaps int, finished *bool) {
	if *finished == true {
		return
	}

	*current++

	if *current >= totalLaps {
		*finished = true
	}

}
