package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	codes := []int{200, 201, 404, 500, 503}
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	for _, v := range codes {
		url := fmt.Sprintf("https://httpbin.org/status/%d", v)

		resp, err := httpClient.Get(url)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			fmt.Printf("%d: успех\n", resp.StatusCode)
		} else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			fmt.Printf("%v: ошибка клиента\n", resp.StatusCode)
		} else if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			fmt.Printf("%v: ошибка сервера\n", resp.StatusCode)
		}
	}

}
