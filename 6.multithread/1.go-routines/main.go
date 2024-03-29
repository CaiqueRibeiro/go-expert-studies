package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s Running!\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("A")
	go task("B")
	time.Sleep(15 * time.Second) // Esse é um jeito de esperar as go-routines terminarem (existem muitos jeitos melhores e built-in da linguagem)
}
