package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	req.Header.Set("X-STUDENT", "Nekdil")

	resp, err := httpClient.Do(req)
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

	fmt.Println(string(data))
}
