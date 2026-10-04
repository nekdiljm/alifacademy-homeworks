package main

import (
	"fmt"
)

func main() {
	name := "Alex"
	company := "SpaceX."
	salary := 23000
	fmt.Println("Меня зовут", name, "и я работаю в", company, "Моя зарплата", salary, "долларов")

	name = "Олег"
	fmt.Println("Меня зовут", name)

	name = "Андрей"
	company = "UAZ"
	salary = 500
	fmt.Println("Меня зовут", name, "и я работаю в", company, "Моя зарплата", salary, "долларов")
}
