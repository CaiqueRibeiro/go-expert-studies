package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(ch)
	go reader(ch, &wg)
	wg.Wait() // Irá esperar todas as goroutines terminarem
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("O número é %d ", x)
		wg.Done() // decrementa o contador de goroutines usando wait group. Com isso não é preciso de um for adicional para não encessar a main
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // preciso fechar o canal para evitar deadlock (forever) na função reader
}
