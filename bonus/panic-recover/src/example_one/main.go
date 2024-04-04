package main

import "fmt"

func panic1() {
	panic("panic 1!")
}

func panic2() {
	panic("panic 2!")
}

func main() {
	defer func() { // function to recover from panic
		if r := recover(); r != nil {
			// Can do a lot of things: send message to telegram, log error, etc
			if r == "panic 1!" {
				fmt.Println("Recovered from panic 1: ", r)
			}
			if r == "panic 2!" {
				fmt.Println("Recovered from panic 2: ", r)
			}
		}
	}()
	panic2()
}
