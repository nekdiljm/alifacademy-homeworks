package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("нельзя извлечь корень из отрицательного числа")
	}

	return math.Sqrt(x), nil
}

func main() {
	sqrtN, err := Sqrt(10)
	if err != nil {
		fmt.Println("Ошибка", err)
	}

	fmt.Printf("Корень от числа: %.3f\n", sqrtN)

	sqrtN, err = Sqrt(-1)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	fmt.Printf("Корень от числа: %.3f\n", sqrtN)

	sqrtN, err = Sqrt(20)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	fmt.Printf("Корень от числа: %.3f\n", sqrtN)

}
