package main

import "fmt"

type Day struct {
	Name  string
	TempC float64
}

func main() {
	days := []Day{
		{Name: "Понедельник", TempC: 24},
		{Name: "Вторник", TempC: 10},
		{Name: "Среда", TempC: 4},
		{Name: "Четверг", TempC: 30},
		{Name: "Пятница", TempC: 45},
		{Name: "Суббота", TempC: 17},
		{Name: "Воскресенье", TempC: 14},
	}

	hot := make(map[string]float64)
	warm := make(map[string]float64)
	cold := make(map[string]float64)

	for _, v := range days {

		switch {
		case v.TempC > 30:
			hot[v.Name] = v.TempC
		case v.TempC >= 15 && v.TempC <= 30:
			warm[v.Name] = v.TempC
		case v.TempC < 15:
			cold[v.Name] = v.TempC
		}
		//if v.TempC > 30 {
		//	hot[v.Name] = v.TempC
		//} else if v.TempC >= 15 && v.TempC <= 30 {
		//	warm[v.Name] = v.TempC
		//} else if v.TempC < 15 {
		//	cold[v.Name] = v.TempC
		//}
	}

	fmt.Printf("Жаркие дни: %v\nТеплые дни: %v\nХолодные дни: %v\n", hot, warm, cold)
}
