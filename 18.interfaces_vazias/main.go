package main

import "fmt"

// INTERFACES VAZIAS
// Usadas para substituir o Generics antes do go 1.18, pois interface sem nada é implementada por tudo

type x interface{}

func main() {

	var x interface{} = 10
	var y interface{} = "Hello, World"
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
