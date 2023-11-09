package main
// PERCORRENDO ARRAYS

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 3
	meuArray[2] = 10

	fmt.Println(len(meuArray)) // imprime o tamanho do array

	// percorrer array
	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é %d\n", i, v)
	}
}