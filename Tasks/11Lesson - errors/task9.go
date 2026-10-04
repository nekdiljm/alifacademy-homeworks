package main

import (
	"fmt"
	"strconv"
)

func ParseAge(s string) (int, error) {
	nums, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Не удалось разобрать возраст: %w", err)
	}

	return nums, nil
}

func main() {
	slice := []string{"25", "abc", "-3", "999"}

	for _, v := range slice {
		num, err := ParseAge(v)
		if err != nil {
			fmt.Println("Ошибка", err)
			continue
		}

		if num < 0 {
			fmt.Printf("Отрицательное значение: %d\n", num)
			continue
		}

		if num > 120 {
			fmt.Printf("Большое число: %d\n", num)
			continue
		}
		fmt.Printf("%d \n", num)
	}
}
