package main

import "fmt"

func main() {
	var (
		name    string
		surname string
		age     uint
		height  float64
		weight  float64
		isWork  bool
	)
	fmt.Println("Здравствуйте, как вас зовут?")
	fmt.Scan(&name)
	fmt.Println("Какая у вас фамилия?")
	fmt.Scan(&surname)
	fmt.Println("Сколько вам лет?")
	fmt.Scan(&age)
	fmt.Println("Какой у вас рост?")
	fmt.Scan(&height)
	fmt.Println("Сколько вы весите?")
	fmt.Scan(&weight)
	fmt.Println("Вы работаете где-либо?(true/false)")
	fmt.Scan(&isWork)

	fmt.Printf("Вы %s, ваша фамилия %s, вам %d лет, ваш рост %.1f, вы весите %.1f. Вы работаете -  %t", name, surname, age, height, weight, isWork)
}
