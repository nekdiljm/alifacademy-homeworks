package main

func main() {
	ch := make(chan int)

	close(ch) // panic: close of closed channel
	close(ch) // Закрытие закрытого канала приводит к панике

	//close(ch) // panic: send on closed channel
	//ch <- 10  // Запись в закрытый канал так же приводит к панике

}
