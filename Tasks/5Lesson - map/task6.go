package main

import "fmt"

func main() {
	m := make(map[string]int, 2)

	m["firstVideo"]++
	m["firstVideo"]++
	m["firstVideo"]++

	m["secondVideo"]++
	m["secondVideo"]++

	fmt.Println(m)

}
