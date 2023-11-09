package main

import "fmt"

// FUNCOES VARIADICAS - recebem parametros de quantidade indeterminada

func main() {
	fmt.Println(sum(1, 2, 3, 7, 45, 7, 3, 6, 3, 6))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
