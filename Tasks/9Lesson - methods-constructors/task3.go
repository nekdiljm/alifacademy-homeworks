package main

import (
	"fmt"
)

type Inventory struct {
	itemCount map[string]int
}

func NewInventory() *Inventory {
	return &Inventory{
		itemCount: make(map[string]int),
	}
}

func (i *Inventory) Add(item string, qty int) {
	if i.itemCount == nil {
		return
	}
	i.itemCount[item] = qty
}

func (i *Inventory) Remove(item string, qty int) error {
	count, ok := i.itemCount[item]
	if !ok {
		return fmt.Errorf("товара нет на складе")
	}

	count -= qty
	if count <= 0 {
		return fmt.Errorf("недостаточное кол-во товара")
	}
	i.itemCount[item] = count

	return nil
}

func (i *Inventory) Count(item string) int {
	return i.itemCount[item]
}

func (i *Inventory) Total() int {
	var totalItem int

	for _, v := range i.itemCount {
		totalItem += v
	}

	return totalItem
}

func main() {
	invent := NewInventory()

	invent.Add("Яблоко", 3)
	invent.Add("Груша", 3)

	if err := invent.Remove("Груша", 2); err != nil {
		fmt.Println("ошибка", err)
		return
	}

	itemCount := invent.Count("Груша")
	fmt.Println("товара", itemCount)

	totalItem := invent.Total()
	fmt.Println("общее кол-во", totalItem)

	fmt.Println(invent)
}
