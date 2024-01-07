package main

// IMPRIMIR TIPOS E FORMATAR STRING

import "fmt"

type ID int

var (
	c int          // infere 0
	d string       // infere ''
	e float64      // infere 0.00000
	f int     = 10 // também é possível atribuir na declaração
	g ID
)

func main() {
	// %T mostra o tipo, enquanto $v mostra o valor da string (%d caso seja dígito)
	fmt.Printf("O Tipo de E é %T", e)
}
