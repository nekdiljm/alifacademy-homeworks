package main

import (
	"errors"
	"fmt"
)

var (
	ErrNegativeAge = errors.New("отрицательный возраст")
	ErrTooOld      = errors.New("слишком высокий возраст")
)

func ValidateAge(age int) error {
	if age < 0 {
		return ErrNegativeAge
	}

	if age > 120 {
		return ErrTooOld
	}

	return nil
}

func main() {
	ages := []int{-5, 20, 30, 150, -3}

	for _, v := range ages {
		if err := ValidateAge(v); err != nil {
			if errors.Is(err, ErrNegativeAge) {
				fmt.Println("Отрицательный возраст")
			}

			if errors.Is(err, ErrTooOld) {
				fmt.Println("Слишком большой возраст")
			}
		}
	}
}
