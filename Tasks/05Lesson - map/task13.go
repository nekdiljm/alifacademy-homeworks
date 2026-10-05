package main

import "fmt"

func main() {
	arr := []string{
		"Душанбе",
		"Москва",
		"Хучанд",
		"Санкт Петербург",
		"Душанбе",
		"Душанбе",
		"Москва",
	}

	m := make(map[string]int)

	for _, v := range arr {
		m[v]++
	}
	fmt.Println(m)

}
