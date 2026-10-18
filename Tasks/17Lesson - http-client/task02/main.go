package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Name: %s\nEmail: %s\nPhone: %s\n", user.Name, user.Email, user.Phone)

}
