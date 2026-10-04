package main

import "fmt"

func main() {
	myTown := "Dushanbe"
	ptr := &myTown

	*ptr = "Moscow"

	fmt.Println(myTown)
	fmt.Println(*ptr)
}
