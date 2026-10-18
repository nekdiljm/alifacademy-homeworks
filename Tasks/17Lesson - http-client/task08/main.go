package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	httpClient := http.Client{
		Timeout: time.Second,
	}

	if _, err := httpClient.Get("https://httpbin.org/delay/2"); err != nil {
		fmt.Println("Error", err)
	}

	httpClient1 := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient1.Get("https://httpbin.org/delay/2")
	if err != nil {
		fmt.Println("Error", err)
	}
	defer resp.Body.Close()
	fmt.Println("StatusCode:", resp.StatusCode)

}
