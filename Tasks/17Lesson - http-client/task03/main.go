package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TitleGet struct {
	Title string `json:"title"`
}

func printTitle(url string) (string, int, error) {
	var title TitleGet

	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", 0, err
		}

		if err := json.Unmarshal(data, &title); err != nil {
			return "", 0, err
		}

	} else if resp.StatusCode == http.StatusNotFound {
		return "Задача не найдена", resp.StatusCode, err
	}

	return title.Title, resp.StatusCode, nil
}

func main() {
	resp, statusCode, err := printTitle("https://jsonplaceholder.typicode.com/todos/5")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	resp2, statusCode1, err := printTitle("https://jsonplaceholder.typicode.com/todos/99999")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("First request answer: %v\nStatusCode: %d\n", resp, statusCode)
	fmt.Printf("Second request answer: %v\nStatusCode: %d\n", resp2, statusCode1)
}
