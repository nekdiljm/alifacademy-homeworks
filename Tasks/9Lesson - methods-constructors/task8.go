package main

import "fmt"

type Student struct {
	Name  string
	Grade []int
}

func (s *Student) Average() float64 {
	if s == nil {
		return -1
	}

	var (
		avg float64
		sum int
	)

	for _, v := range s.Grade {
		sum += v
	}

	avg = float64(sum / len(s.Grade))
	return avg
}

func main() {
	strudents := []Student{
		{
			Name:  "Nekruz",
			Grade: []int{10, 8, 9, 10},
		},
		{
			Name:  "Nekdil",
			Grade: []int{7, 5, 7, 10},
		},
		{
			Name:  "Behzod",
			Grade: []int{3, 2, 6, 10},
		},
		{
			Name:  "Gundil",
			Grade: []int{8, 4, 5, 7},
		},
		{
			Name:  "Nekdul",
			Grade: []int{5, 6, 3, 5},
		},
	}

	var bestAvg float64

	for _, v := range strudents {
		if bestAvg < v.Average() {
			bestAvg = v.Average()
		}
	}

	fmt.Println(bestAvg)

}
