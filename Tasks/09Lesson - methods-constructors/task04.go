package main

import (
	"fmt"
)

type Matrix struct {
	arr [][]float64
}

func (m *Matrix) Rows() int {
	return len(m.arr)
}

func (m *Matrix) Cols() int {
	if len(m.arr) < 0 {
		return -1
	}
	return len(m.arr[0])
}

func (m *Matrix) Sum() float64 {
	var sum float64

	for i := 0; i < len(m.arr); i++ {
		for j := 0; j < len(m.arr[i]); j++ {
			sum += m.arr[i][j]
		}
	}

	return sum
}

func (m *Matrix) Transpose() Matrix {
	return Matrix{}
}

func main() {
	m := Matrix{
		arr: [][]float64{
			{
				1, 2, 3,
			},
			{
				4, 5, 6,
			},
		},
	}

	fmt.Println(m.Cols())
	fmt.Println(m.Rows())
	fmt.Println(m.Sum())

}
