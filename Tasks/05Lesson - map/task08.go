package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"Некдил":    4,
		"Самир":     3,
		"Тимур":     5,
		"Абдурахим": 2,
	}

	arr := make([]string, 0, len(m))

	for k := range m {
		arr = append(arr, k)
	}

	sort.Strings(arr)

	for _, k := range arr {
		fmt.Println(k, m[k])
	}

}
