package main

import (
	"errors"
	"fmt"
)

func step1(n int) error {
	if n < 0 {
		return errors.New("отрицательное число")
	}

	return nil
}

func step2(n int) error {
	if n%3 == 0 {
		return errors.New("число делится на 3")
	}

	return nil
}

func step3(n int) error {
	if n%2 == 0 {
		return errors.New("число делится на 2")
	}

	return nil
}

func Pipeline(n int) error {
	if err := step1(n); err != nil {
		return err
	}

	if err := step2(n); err != nil {
		return err
	}

	if err := step3(n); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Pipeline(3); err != nil {
		fmt.Println("Ошибка", err)
	}
}
