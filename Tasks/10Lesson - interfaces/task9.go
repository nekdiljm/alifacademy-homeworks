package main

import "fmt"

type Box struct {
	Value any
}

type Test struct {
	test1 string
}

func main() {
	boxes := []Box{
		{Value: "String"},
		{Value: 14},
		{Value: 23.3},
		{Value: false},
		{Value: []int{}},
		{Value: map[string]int{}},
		{Value: Test{}},
	}

	for _, v := range boxes {
		fmt.Printf("%T %v\n", v.Value, v.Value)
	}
}
