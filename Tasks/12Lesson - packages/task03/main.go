package main

import (
	"fmt"
	"task3/validator"
)

func main() {
	email := "test@gmail.com"
	age := 18
	isValidEmail := validator.IsValidEmail(email)
	isAdult := validator.IsAdult(age)
	fmt.Printf("isValidEmail: %t || isAdult: %t\n", isValidEmail, isAdult)

	email = "testgmail.com"
	age = 15
	isValidEmail = validator.IsValidEmail(email)
	isAdult = validator.IsAdult(age)
	fmt.Printf("isValidEmail: %t || isAdult: %t\n", isValidEmail, isAdult)

	email = "test@gmail.com"
	age = 12
	isValidEmail = validator.IsValidEmail(email)
	isAdult = validator.IsAdult(age)
	fmt.Printf("isValidEmail: %t || isAdult: %t\n", isValidEmail, isAdult)

	email = "test@gmailcom"
	age = 21
	isValidEmail = validator.IsValidEmail(email)
	isAdult = validator.IsAdult(age)
	fmt.Printf("isValidEmail: %t || isAdult: %t\n", isValidEmail, isAdult)
}
