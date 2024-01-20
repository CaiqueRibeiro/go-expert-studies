package main

import (
	"fmt"
	"sync"
	"time"
)

// wait groups são métodos para que seja possível esperar que um conjunto de goroutines termine de executar.

func task(name string, w *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s Running!\n", i, name)
		time.Sleep(1 * time.Second)
		w.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)

	go task("A", &waitGroup)
	go task("B", &waitGroup)

	go func(name string, w *sync.WaitGroup) {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s Running!\n", i, name)
			time.Sleep(1 * time.Second)
			w.Done()
		}
	}("anonymous", &waitGroup)

	waitGroup.Wait()
}