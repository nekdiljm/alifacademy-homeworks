package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient.Get("https://jsonplaceholder.typicode.com/todos")
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

	var todo []Todo
	if err := json.Unmarshal(data, &todo); err != nil {
		fmt.Println("Error", err)
		return
	}
	var (
		completedTask   int
		uncompletedTask int
	)
	for _, v := range todo {
		if v.Completed {
			completedTask++
		} else {
			uncompletedTask++
		}
	}

	fmt.Printf("Выполненные задачи: %d\nНевыполненные задачи: %d\n", completedTask, uncompletedTask)

}
