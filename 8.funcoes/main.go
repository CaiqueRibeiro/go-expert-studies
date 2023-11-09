package main

import (
	"errors"
	"fmt"
)

// FUNCOES

func main() {
	valor, err := sum3(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a, b int) int { // se tiverem o mesmo tipo, pode ser passado assim
	return a + b
}

func sum3(a, b int) (int, error) { // é possivel retornar 2 valores
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
