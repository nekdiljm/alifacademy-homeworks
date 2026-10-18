package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	params := url.Values{}
	params.Set("userId", "3")
	query := params.Encode()

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient.Get("https://jsonplaceholder.typicode.com/posts?" + query)
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

	fmt.Println("Кол-во постов:", len(posts))
	for _, v := range posts {
		fmt.Println("title:", v.Title)
	}

}
