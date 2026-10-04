package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName, lastName := parseFullName("Джамшедзод Некдил")
	fmt.Println(firstName, lastName)
}

func parseFullName(full string) (string, string) {
	parts := strings.SplitN(full, " ", 2)

	return parts[0], parts[1]
}
