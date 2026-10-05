package main

import "fmt"

func main() {
	avgMark := averageGrade(map[string]int{
		"Коля": 5,
		"Вася": 4,
		"Гена": 3,
		"Петя": 4,
	})
	fmt.Printf("Средний балл учеников: %.1f", avgMark)
}

func averageGrade(grades map[string]int) float64 {
	var (
		avgMark int
		count   = len(grades)
	)

	for _, v := range grades {
		avgMark = (avgMark + v)
	}
	return float64(avgMark) / float64(count)
}
