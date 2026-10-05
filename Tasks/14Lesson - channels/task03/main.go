package main

func main() {
	ch := make(chan int)

	ch <- 10
}

// fatal error: all goroutines are asleep - deadlock!
// происходит потому-что мы отправляем в канал значение, но никто его не читает и программа бесконечно ждет, пока кто-то прочитает
