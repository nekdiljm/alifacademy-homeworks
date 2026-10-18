package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
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

	fmt.Println("Статус код:", resp.StatusCode)
	fmt.Println("Content type:", resp.Header.Get("Content-Type"))
	fmt.Println("Тело ответа:", string(data))
}
