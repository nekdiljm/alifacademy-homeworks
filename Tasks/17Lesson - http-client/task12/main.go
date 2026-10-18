package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type company struct {
	Name string `json:"name"`
}
type User struct {
	Name    string  `json:"name"`
	Company company `json:"company"`
}

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected statusCode:", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	users := []User{}
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println("Error", err)
		return
	}

	for _, v := range users {
		fmt.Printf("name: %s, companyName: %s\n", v.Name, v.Company.Name)
	}
}
