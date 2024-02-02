package main

import "fmt"

func main() {
	ch := make(chan int)

	go publish(ch)
	reader(ch)
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("O número é %d ", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // preciso fechar o canal para evitar deadlock (forever) na função reader
}
