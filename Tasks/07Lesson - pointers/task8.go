package main

import "fmt"

func main() {
	logInCity := map[string]string{
		"Alex":    "Moscow",
		"John":    "New York",
		"Umed":    "Dushanbe",
		"Ulugbek": "Tashkent",
	}
	someLogin := "Alex"
	cityCheck := findCity(logInCity, someLogin)

	if cityCheck != nil {
		fmt.Printf("Логин: %v, город: %v", someLogin, *cityCheck)
	} else {
		fmt.Printf("Неверный логин")
	}

}

func findCity(users map[string]string, login string) *string {
	city, ok := users[login]
	if !ok {
		return nil
	}
	return &city
}
