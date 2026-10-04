package main

import (
	"fmt"
	"task2/strutil"
)

func main() {
	word := "дед"
	reverse := strutil.Reverse(word)
	isPalindrome := strutil.IsPalindrome(reverse)
	fmt.Printf("Слово до реверса: %s || Слово после: %s\nПолиндром: %t\n", word, reverse, isPalindrome)

	word = "такси"
	reverse = strutil.Reverse(word)
	isPalindrome = strutil.IsPalindrome(reverse)
	fmt.Printf("Слово до реверса: %s || Слово после: %s\nПолиндром: %t\n", word, reverse, isPalindrome)

	word = "коза на мази"
	reverse = strutil.Reverse(word)
	isPalindrome = strutil.IsPalindrome(reverse)
	fmt.Printf("Слово до реверса: %s || Слово после: %s\nПолиндром: %t\n", word, reverse, isPalindrome)
}
