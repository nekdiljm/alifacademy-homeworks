package main

import "fmt"

func main() {
	secrNumb := 37
	var userNum int

	for {
		fmt.Print("Введите число - ")
		fmt.Scan(&userNum)

		if userNum < secrNumb {
			fmt.Printf("Число больше!\n\n")
		} else if userNum > secrNumb {
			fmt.Println("Число меньше!\n")
		} else if userNum == secrNumb {
			break
		}
	}
}
