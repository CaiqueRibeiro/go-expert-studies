package main

import "fmt"

// TYPE ASSERTIONS
// Mesmo passando interfaces vazias, é possivel forçar o tipo caso tenha certeza deste tipo

type x interface{}

func main() {
	var minhaVar interface{} = "Caique Ribeiro"
	println(minhaVar)          // irá imprimir coisas malucas pois não sabe o tipo
	println(minhaVar.(string)) // type assertion para garantir que seja string
	res, ok := minhaVar.(int)  // para garantir que não dê exception em assertions erradas
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}
