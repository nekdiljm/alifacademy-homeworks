package main

import "fmt"

func main() {
	div, rem := divmod(234, 5)
	fmt.Printf(("Частное a и b: %d\nОстаток: %d"), div, rem)
}

func divmod(a, b int) (int, int) {
	div, rem := a/b, a%b
	return div, rem
}
