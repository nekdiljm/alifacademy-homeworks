package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	Time        time.Time
}

func (p Payment) Println() {
	fmt.Println("Description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Address:", p.Address)
}

var (
	money          = 1000
	paymentHistory = make([]Payment, 0)
	mu             sync.Mutex
)

type HttpResponse struct {
	Money          int
	PaymentHistory []Payment
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payment.Time = time.Now()

	mu.Lock()
	defer mu.Unlock()
	if money-payment.USD >= 0 {
		money -= payment.USD
	} else {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Недостаточный баланс"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("Error", err)
			return
		}
		return
	}

	paymentHistory = append(paymentHistory, payment)
	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	data, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := w.Write(data); err != nil {
		fmt.Println("Error", err)
		return
	}
}

func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
		return
	}
}
