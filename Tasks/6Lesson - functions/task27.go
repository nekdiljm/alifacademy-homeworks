package main

import (
	"fmt"
	"unicode"
)

func main() {
	lettersInStr, diggitsInStr := countLetterAndDigits("Привет23")
	fmt.Printf("Кол-во букв в строке: %d\nКол-во цифр в строке: %d", lettersInStr, diggitsInStr)
}

func countLetterAndDigits(s string) (int, int) {
	var (
		letter int
		digits int
	)

	for _, v := range s {
		if unicode.IsLetter(v) {
			letter++
		} else if unicode.IsDigit(v) {
			digits++
		}
	}
	return letter, digits
}
