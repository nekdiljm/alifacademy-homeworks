package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient.Get("https://jsonplaceholder.typicode.com/posts")
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

	var posts []Post
	if err := json.Unmarshal(data, &posts); err != nil {
		fmt.Println("Error", err)
		return
	}

	lastEl := len(posts) - 1
	fmt.Println("Кол-во постов:", len(posts))
	fmt.Printf("title первого: %s\ntitle последнего: %s\n", posts[0].Title, posts[lastEl].Title)
}
