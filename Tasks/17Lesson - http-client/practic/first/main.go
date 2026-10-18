package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var (
	mu      sync.Mutex
	balance = 50
	bank    = 0
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "Error HTTP request body" + err.Error()

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
			return
		}
		return
	}

	httpRequestBodyStr := string(httpRequestBody)

	payAmount, err := strconv.Atoi(httpRequestBodyStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Error to convert str to int " + err.Error()

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if balance-payAmount >= 0 && payAmount > 0 {
		balance -= payAmount

		msg := "Оплата прошла успешно. Баланс:" + strconv.Itoa(balance)

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}
	} else {
		msg := "Error, no enough funds to pay or negative num"
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}
		return
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "Error, HTTP request body" + err.Error()

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}
		return
	}

	httpRequestBodyStr := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyStr)
	if err != nil {
		msg := "Error to convert str to int" + err.Error()

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if balance-saveAmount >= 0 && saveAmount > 0 {
		balance -= saveAmount
		bank += saveAmount

		msg := "Bank successfully saved. Bank:" + strconv.Itoa(bank) + " || Баланс:" + strconv.Itoa(balance)

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}
	} else {
		msg := "Error, no enough funds to save or negative num"

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}
		return
	}
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "Error HTTP request body" + err.Error()

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
			return
		}
		return
	}

	httpRequestBodyStr := string(httpRequestBody)

	depositAmount, err := strconv.Atoi(httpRequestBodyStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Error to convert str to int" + err.Error()

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if depositAmount > 0 {
		balance += depositAmount

		msg := "Баланс пополнен. Баланс:" + strconv.Itoa(balance)

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Error negative num"

		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error to write HTTP response", err)
		}
	}

}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/deposit", depositHandler)

	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("Error to start listen and serve", err)
		return
	}

}
