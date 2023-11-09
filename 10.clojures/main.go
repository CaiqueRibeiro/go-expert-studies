package main

import "fmt"

// CLOJURES (ou funções anônimas)

func main() {
	total := func() int {
		return sum(1, 5, 7, 3, 3) * 2
	}() // Função anônima (que pode ser uma clojure)

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
