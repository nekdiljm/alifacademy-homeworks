package main

import "fmt"

const post = "Go Backend Developer"

func main() {
	var (
		name        string
		surname     string
		age         uint
		town        string
		exp         float64
		isEducation bool
		telegramUs  string
	)
	fmt.Println("Как вас зовут?")
	fmt.Scan(&name)
	fmt.Println("Какая у вас фамилия?")
	fmt.Scan(&surname)
	fmt.Println("Сколько вам лет?")
	fmt.Scan(&age)
	fmt.Println("Где вы живете?")
	fmt.Scan(&town)
	fmt.Println("Какой у вас опыт работы?(число)")
	fmt.Scan(&exp)
	fmt.Println("У вас есть образование?(true/false)")
	fmt.Scan(&isEducation)
	fmt.Println("Можете написать свой telegram")
	fmt.Scan(&telegramUs)

	fmt.Printf("Меня зовут %s. Моя фамилия %s\nМне %d лет\nЯ живу в %s\nЯ подаю на должность %v\nМой опыт работы %.1f года\nМое образование - %t\nМой телеграм %s", name, surname, age, town, post, exp, isEducation, telegramUs)
}
