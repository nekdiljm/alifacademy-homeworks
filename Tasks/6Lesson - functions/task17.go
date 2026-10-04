package main

import "fmt"

func main() {
	slLen := sliceLenght([]string{"Hello", "World!"})
	fmt.Printf("Длина слайса: %v", slLen)
}

func sliceLenght(s []string) int {
	return len(s)
}
