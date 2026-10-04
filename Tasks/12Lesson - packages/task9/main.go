package main

import (
	"fmt"
	"task9/user"
)

func main() {
	user1 := user.NewUser("Nekdil")
	fmt.Printf("Name: %s || ID: %v\n", user1.Name, user1.ID)

	user2 := user.NewUser("Nekruz")
	fmt.Printf("Name: %s || ID: %v\n", user2.Name, user2.ID)

	user3 := user.NewUser("Behzod")
	fmt.Printf("Name: %s || ID: %v\n", user3.Name, user3.ID)

}
