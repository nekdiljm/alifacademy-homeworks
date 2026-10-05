package main

import "fmt"

func main() {
	var role string

	fmt.Print("Введите роль(admin, manager, user, guest) - ")
	fmt.Scan(&role)

	if role == "admin" {
		fmt.Println("У вас все права доступа!")
	} else if role == "manager" {
		fmt.Println("У вас права управления пользователями!")
	} else if role == "user" {
		fmt.Println("У вас обычный доступ!")
	} else if role == "guest" {
		fmt.Println("Вам доступен только просмотр!")
	} else {
		fmt.Println("Такой роли нет! Выберите подходящую!")
		return
	}
}
