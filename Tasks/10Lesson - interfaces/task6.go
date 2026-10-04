package main

import "fmt"

type Transport interface {
	Price(distanceKm float64) float64
	Name() string
}

type Taxi struct {
	landing float64
	priceKm float64
	name    string
}

type Bus struct {
	landing float64
	name    string
}

type Bike struct {
	priceKm float64
	name    string
}

func (t Taxi) Price(distanceKm float64) float64 {
	return t.landing + (t.priceKm * distanceKm)
}

func (t Taxi) Name() string {
	return t.name
}

func (b Bus) Price(_ float64) float64 {
	return b.landing
}

func (b Bus) Name() string {
	return b.name
}

func (bk Bike) Price(distanceKm float64) float64 {
	return bk.priceKm * distanceKm
}

func (bk Bike) Name() string {
	return bk.name
}

func main() {
	transport := []Transport{
		Taxi{landing: 10, priceKm: 4, name: "Jura"},
		Bus{landing: 3, name: "Electrobus"},
		Bike{priceKm: 2, name: "Yamaha"},
	}

	cheaper := transport[0]
	for _, v := range transport {
		if cheaper.Price(5) > v.Price(5) {
			cheaper = v
		}
	}

	fmt.Printf("Самая дешевая цена: %v, у %s\n", cheaper.Price(5), cheaper.Name())
}
