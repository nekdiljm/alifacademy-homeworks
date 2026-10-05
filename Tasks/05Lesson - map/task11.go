package main

import "fmt"

func main() {
	m := map[string]int{
		"Али":    1000,
		"Собир":  2500,
		"Олим":   3500,
		"Чамшед": 7000,
	}

	var (
		maxSalary int
		name      string
	)

	for k, v := range m {

		if maxSalary < v {
			maxSalary = v
			name = k
		}

	}
	fmt.Println(name, maxSalary)
}
