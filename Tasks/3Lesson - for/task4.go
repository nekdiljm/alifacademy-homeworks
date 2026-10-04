package main

import "fmt"

func main() {

	var (
		password = 1234
		usAnsw   int
	)

	for i := 1; i <= 3; i++ {
		fmt.Print("Введите пароль - ")
		fmt.Scan(&usAnsw)
		if usAnsw == password {
			fmt.Println("Доступ разрешен!")
			break
		}

	}
	if usAnsw != password {
		fmt.Println("Доступ запрещен!")
	}
}
