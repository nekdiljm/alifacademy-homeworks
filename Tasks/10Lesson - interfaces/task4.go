package main

import "fmt"

type Instrument interface {
	Play() string
}

type Guitar struct {
	Name string
}

type Piano struct {
	Name string
}

type Drum struct {
	Name string
}

func (g Guitar) Play() string {
	return fmt.Sprintf("Guitar %s: дзынь-дзынь", g.Name)
}

func (p Piano) Play() string {
	return fmt.Sprintf("Piano %s: динь-динь", p.Name)
}

func (d Drum) Play() string {
	return fmt.Sprintf("Drum %s: бум-бум", d.Name)
}

func main() {
	instruments := []Instrument{
		Guitar{Name: "Fender"},
		Piano{Name: "Yamaha"},
		Drum{Name: "Pearl"},
		Guitar{Name: "Gibson"},
		Piano{Name: "Casio"},
	}

	for _, v := range instruments {
		fmt.Println(v.Play())
	}
}
