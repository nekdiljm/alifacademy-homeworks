package main

import (
	"fmt"
)

func main() {
	a := fullName("Nekdil", "Jamshedzod")
	fmt.Println(a)
}

func fullName(firstName, lastName string) string {
	glueString := firstName + " " + lastName
	return glueString

}
