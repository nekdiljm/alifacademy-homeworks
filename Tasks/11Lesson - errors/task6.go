package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyLogin    = errors.New("логин пуст")
	ErrShortPassword = errors.New("пароль короче 6 символов")
)

func CheckCredentials(login, password string) error {
	if login == "" {
		return ErrEmptyLogin
	}

	if len(password) < 6 {
		return ErrShortPassword
	}

	return nil
}

func main() {
	err := CheckCredentials("Nekdil", "123456")
	if err != nil {
		switch {
		case errors.Is(err, ErrEmptyLogin):
			fmt.Println("Ошибка: пустой логин")
			return
		case errors.Is(err, ErrShortPassword):
			fmt.Println("Ошибка: пароль короче 6 симолов")
			return
		}

	}

	fmt.Println("Доступ разрешен!")

}
