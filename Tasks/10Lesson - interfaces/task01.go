package main

import "fmt"

type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct {
	Name string
}

type Bus struct {
	Name string
}

type Bike struct {
	Name string
}

func (c Car) Start() string {
	return fmt.Sprintf("Car %s заводится\n", c.Name)
}

func (c Car) Stop() string {
	return fmt.Sprintf("Car %s глохнет\n", c.Name)
}

func (b Bus) Start() string {
	return fmt.Sprintf("Bus %s заводится\n", b.Name)
}

func (b Bus) Stop() string {
	return fmt.Sprintf("Bus %s глохнет\n", b.Name)
}

func (bk Bike) Start() string {
	return fmt.Sprintf("Bike %s заводится\n", bk.Name)
}

func (bk Bike) Stop() string {
	return fmt.Sprintf("Bike %s глохнет\n", bk.Name)
}
func main() {
	vehicles := []Vehicle{
		Car{Name: "Toyota"},
		Bus{Name: "ПАЗ"},
		Bike{Name: "Yamaha"},
	}

	for _, v := range vehicles {
		fmt.Print(v.Start())
		fmt.Print(v.Stop())
	}
}
