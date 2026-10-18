package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	user := User{
		Name:  "Nekruz",
		Age:   26,
		Email: "Email",
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(data))
}
