package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("StatusCode:", resp.StatusCode)
}
