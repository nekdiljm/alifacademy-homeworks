package main

import "fmt"

func main() {
	var (
		cardBal  int
		cashTake int
	)
	fmt.Print("Введите баланс карты: ")
	fmt.Scan(&cardBal)

	fmt.Print("Введите сумму снятия: ")
	fmt.Scan(&cashTake)

	if cashTake <= 0 || cardBal <= 0 {
		fmt.Println("Ошибка, вы ввели неверный баланс или сумму снятия!")
	} else if cashTake > cardBal {
		fmt.Println("Недостаточно средст!")
	} else {
		remains := cardBal - cashTake
		fmt.Println("Остаток средств на карте = ", remains)
	}

	//switch {
	//case cashTake <= 0 || cardBal <= 0:
	//	fmt.Println("Ошибка, вы ввели неверный баланс или сумму снятия!")
	//case cashTake > cardBal:
	//	fmt.Println("Недостаточно средст!")
	//default:
	//	remains := cardBal - cashTake
	//	fmt.Println("Остаток средств на карте = ", remains)
	//}

}
