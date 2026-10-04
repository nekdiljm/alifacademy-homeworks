package main

import (
	"errors"
	"fmt"
)

type Order struct {
	ID     int
	Amount float64
}

type OrderQueue struct {
	order []Order
}

func (oq *OrderQueue) Enqueue(o Order) {
	if oq == nil {
		return
	}

	oq.order = append(oq.order, o)
}

func (oq *OrderQueue) Dequeue() (Order, error) {
	if oq == nil {
		return Order{}, errors.New("пустой ресивер")
	}

	if oq.IsEmpty() {
		return Order{}, errors.New("пустая очередь")
	}

	firstIdx := 0
	firstEl := oq.order[firstIdx]

	if firstIdx == len(oq.order)-1 {
		oq.order = oq.order[:firstIdx]
	} else {
		oq.order = oq.order[firstIdx+1:]
	}

	return firstEl, nil

	// LIFO
	//lastIdx := len(oq.order) - 1
	//lastEl := oq.order[lastIdx]
	//oq.order = oq.order[:lastIdx]
	//return lastEl, nil
}

func (oq *OrderQueue) IsEmpty() bool {
	if oq == nil {
		return true
	}

	return len(oq.order) == 0
}

func (oq *OrderQueue) TotalAmount() float64 {
	if oq == nil {
		return -1
	}

	var total float64

	for _, v := range oq.order {
		total += v.Amount
	}

	return total
}

func main() {
	orders := OrderQueue{}

	orders.Enqueue(Order{
		ID:     1,
		Amount: 30,
	})

	orders.Enqueue(Order{
		ID:     2,
		Amount: 30,
	})

	orders.Enqueue(Order{
		ID:     3,
		Amount: 30,
	})

	ord, err := orders.Dequeue()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(ord)

	ord, err = orders.Dequeue()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(ord)

	ord, err = orders.Dequeue()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(ord)

	ord, err = orders.Dequeue()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(ord)
	total := orders.TotalAmount()
	fmt.Println(total)

}
