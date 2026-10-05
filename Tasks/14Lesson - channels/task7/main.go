package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	v, ok := <-ch
	fmt.Println(v, ok)

	v, ok = <-ch
	fmt.Println(v, ok)

	v, ok = <-ch
	fmt.Println(v, ok)

	v, ok = <-ch
	fmt.Println(v, ok)
}

// Чтение из закрытого пустого канала возвращает default value типа
// ok == true - значит значение из канала успешно получено
// ok == false - значите канал пустой и читаем деволтное значение
