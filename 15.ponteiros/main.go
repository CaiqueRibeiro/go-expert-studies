package main

import "fmt"

// PONTEIROS

func main() {
	a := 10
	fmt.Println(a)                                    // mostra 10
	fmt.Println(&a)                                   // mostra o endereço de memória onde esta o 10
	var ponteiro *int = &a                            // o ponteiro guarda o endereço de memória
	fmt.Println(ponteiro)                             // retorna o mesmo endereço de &a
	fmt.Println(&ponteiro)                            // aponta para o endereço onde está a variavel "ponteiro"
	*ponteiro = 20                                    // como o ponteiro está direto no endereço memória, altera o valor
	fmt.Printf("Variavel a: %d\n", a)                 // irá mostrar 20
	fmt.Printf("Variavel *ponteiro: %d\n", *ponteiro) // irá mostrar 20

	fmt.Println("\nRESUMINDO: ")
	fmt.Printf("a: %d\n", a)                 // irá mostrar 20
	fmt.Printf("&a: %d\n", &a)               // irá mostrar endereço X
	fmt.Printf("ponteiro: %d\n", ponteiro)   // irá mostrar o endereço X
	fmt.Printf("*ponteiro: %d\n", *ponteiro) // irá mostrar 20 (o mesmo de a)
	fmt.Printf("&ponteiro: %d\n", &ponteiro) // irá mostrar o endereço da propria variavel ponteiro
}
