package main

import "fmt"

type Warrior interface {
	Attack() int
	Description() string
}

type Knight struct {
	Name     string
	Strength int
}

type Archer struct {
	Name      string
	Precision int
}

type Mage struct {
	Name      string
	ManaPower int
}

func (k Knight) Attack() int {
	return k.Strength * 2
}

func (k Knight) Description() string {
	return "Мечник " + k.Name
}

func (a Archer) Attack() int {
	return a.Precision + 10
}

func (a Archer) Description() string {
	return "Лучник " + a.Name
}

func (m Mage) Attack() int {
	return m.ManaPower * 3
}

func (m Mage) Description() string {
	return "Маг " + m.Name
}

func main() {
	warriors := []Warrior{
		Knight{Name: "Artur", Strength: 25},
		Archer{Name: "Hoyka", Precision: 30},
		Mage{Name: "Marlin", ManaPower: 15},
	}

	var sumAttack int
	for _, v := range warriors {
		fmt.Printf("%s %d\n", v.Description(), v.Attack())
		sumAttack += v.Attack()
	}
	fmt.Println("Суммарная сила атаки отряда: ", sumAttack)
}
