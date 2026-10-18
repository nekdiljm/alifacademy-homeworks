package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string `json:"name"`
	Grade   int    `json:"grade"`
	Comment string `json:"comment,omitempty"`
}

func main() {
	studWithCom := Student{
		Name:    "Ali",
		Grade:   4,
		Comment: "коммент",
	}

	studWithOutCom := Student{
		Name:  "Umed",
		Grade: 5,
	}

	stud1, err := json.Marshal(studWithCom)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	stud2, err := json.Marshal(studWithOutCom)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(stud1))
	fmt.Println(string(stud2))
}
