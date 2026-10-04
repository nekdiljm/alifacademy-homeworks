package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var password string

	fmt.Print("Введите пароль не менее 8 символов: ")
	fmt.Scan(&password)
	ok, pass := checkPassword(password)
	fmt.Println(ok, pass)
}

func checkPassword(pass string) (bool, string) {
	if utf8.RuneCountInString(pass) >= 8 {
		return true, "OK"
	} else {
		return false, "Слишком короткий пароль"
	}
}
