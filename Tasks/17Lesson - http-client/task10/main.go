package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

type UserWithId struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	post := User{
		Title:  "First post",
		Body:   "Json from go",
		UserId: 1,
	}

	data, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	resp, err := httpClient.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer resp.Body.Close()
	
	user := UserWithId{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(user.ID)

}
