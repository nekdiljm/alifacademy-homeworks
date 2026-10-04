package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go это просто go это быстро go это надёжно"

	words := strings.Fields(text)
	m := make(map[string]int)

	for _, v := range words {
		m[v]++
	}
	fmt.Println("Частота каждого слова:", m)

	var (
		maxWord   int
		foundWord string
	)

	for k, v := range m {

		if v > maxWord {
			maxWord = v
			foundWord = k
		}
	}
	fmt.Printf("Самое частое слово - %s:%d", foundWord, maxWord)
}
