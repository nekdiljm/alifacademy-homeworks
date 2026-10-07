package main

import "fmt"

func main() {
	user1 := "Misha Lox" // 0x00001
	user2 := "Maksim"    // 0x0002
	fmt.Printf("До свапа: user1: %v, user2: %v\n\n", user1, user2)

	swapStrings(&user1, &user2)

	fmt.Printf("\nПосле свапа: user1: %v, user2: %v\n", user1, user2)
}

func swapStrings(a, b *string) {
	a, b = b, a
	fmt.Println(a, b)
}
