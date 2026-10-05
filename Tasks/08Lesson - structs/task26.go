package main

import "fmt"

type Student struct {
	Name  string
	Grade int
}

func main() {
	students := []Student{
		{Name: "Nekruz", Grade: 100},
		{Name: "Behruz", Grade: 95},
		{Name: "Ulugbek", Grade: 10},
		{Name: "Nekdil", Grade: 60},
		{Name: "Behzod", Grade: 83},
		{Name: "Dili", Grade: 7},
	}

	m := make(map[string][]string)

	for _, v := range students {
		switch {
		case v.Grade >= 90:
			m["A"] = append(m["A"], v.Name)
		case v.Grade >= 75 && v.Grade <= 89:
			m["B"] = append(m["B"], v.Name)
		case v.Grade >= 60 && v.Grade <= 74:
			m["C"] = append(m["C"], v.Name)
		case v.Grade < 60:
			m["F"] = append(m["F"], v.Name)
		}
	}

	for k, v := range m {
		fmt.Printf("Имя: %v, Оценка: %v\n", m[k], k)
		fmt.Printf("Кол-во оценок %v: %v\n", k, len(v))
	}
}
