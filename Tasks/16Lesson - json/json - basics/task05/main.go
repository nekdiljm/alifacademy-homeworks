package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Login    string  `json:"login"`
	Password string  `json:"-"`
	Balance  float64 `json:"balance"`
}

func main() {
	account := Account{
		Login:    "Alex",
		Password: "alex2308",
		Balance:  1500,
	}

	data, err := json.Marshal(account)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(string(data))
}
