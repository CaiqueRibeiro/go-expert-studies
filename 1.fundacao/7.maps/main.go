package main
// MAPS -> Equivalente ao HashTable

import "fmt"

func main() {
	salarios := map[string]int{"Caique": 9000, "Zezinho": 3900}

	// remover por chave
	delete(salarios, "Caique")

	// adicionar
	salarios["Armando"] = 5000

	// Usando o make -> só cria o mapa, sem preenche-lo
	sal := make(map[string]int)
	fmt.Println(sal)

	// também pode ser declarado assim:
	sal2 := map[string]int{}
	fmt.Println(sal2)

	// percorrendo o map
	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}
}