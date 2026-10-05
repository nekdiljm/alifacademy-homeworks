package main

import "fmt"

type Passenger struct {
	Name string
	Age  int
}

func main() {
	passengers := []Passenger{
		{Name: "Ulugbek", Age: 62},
		{Name: "Alijon", Age: 9},
		{Name: "Aminjon", Age: 7},
		{Name: "Nekruz", Age: 75},
		{Name: "Umed", Age: 35},
		{Name: "Nekdil", Age: 17},
	}

	var (
		children   int
		adult      int
		pensioners int
		totalCash  int
	)

	for _, v := range passengers {
		if v.Age < 12 {
			children++
		} else if v.Age <= 64 {
			adult++
			totalCash += 50
		} else {
			pensioners++
			totalCash += 25
		}

		//switch {
		//case v.Age < 12:
		//	children++
		//case v.Age <= 64:
		//	adult++
		//	totalCash += 50
		//default:
		//	pensioners++
		//	totalCash += 25
		//}
	}
	allPassengers := children + adult + pensioners
	fmt.Printf("Кол-во пассажиров: %d\nДетей: %d Взрослых: %d Пенсиронеров: %d\n", allPassengers, children, adult, pensioners)
	fmt.Printf("Общая выручка: %d сомони\n", totalCash)
}
