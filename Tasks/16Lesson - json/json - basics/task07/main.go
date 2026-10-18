package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	input := `{"name":"Алия","age":20,"city":"Душанбе"}`

	var person Person
	if err := json.Unmarshal([]byte(input), &person); err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(person)
}

// {Алия 20}
// city просто проигнорировался, т.к в структуре нет такого поля
