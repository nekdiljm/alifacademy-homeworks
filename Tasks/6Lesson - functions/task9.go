package main

import "fmt"

func main() {
	seconds := 3600
	secToMin := secondToMinutes(seconds)
	fmt.Printf("%d секунд = %d мин.", seconds, secToMin)
}

func secondToMinutes(second int) int {
	secToMin := second / 60
	return secToMin
}
