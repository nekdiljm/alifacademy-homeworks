package main

import "fmt"

func main() {
	dedup := wordFrequency([]string{"Привет", "Привет", "Hello", "Paris", "Paris", "Paris"})
	fmt.Println(dedup)
}

func wordFrequency(words []string) map[string]int {
	m := make(map[string]int)

	for _, v := range words {
		m[v]++
	}
	return m
}
