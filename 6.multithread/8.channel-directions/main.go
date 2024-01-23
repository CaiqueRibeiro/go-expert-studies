package main

import "fmt"

func recebe(ch chan<- string, hello string) {
	ch <- hello
}

func ler(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string)
	go recebe(ch, "Hello")
	ler(ch)
}
