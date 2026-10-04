package main

import "fmt"

func Divide(a, b float64) (float64, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("divide negative nums(%.1f, %.1f)", a, b)
	}

	if b == 0 {
		return 0, fmt.Errorf("divide by zero %.1f and %.1f", a, b)
	}

	divide := a / b
	return divide, nil
}

func main() {
	divide, err := Divide(10, 0)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Result:", divide)
	fmt.Println("")

	divide, err = Divide(10, 20)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Result:", divide)
	fmt.Println("")

	divide, err = Divide(-1, 20)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Result:", divide)
}
