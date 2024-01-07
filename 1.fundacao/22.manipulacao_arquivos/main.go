package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// // criar arquivo
	// f, err := os.Create("arquivo.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// // escrever arquivo
	// tamanho, err := f.WriteString("Hello, World!") // pode ser apenas Write([]byte("Hello, World")) se for em bytes
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho)
	// defer f.Close()

	// // ler arquivo
	// arquivo, err := os.ReadFile("arquivo.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(arquivo))

	// ler arquivo em stream (arquivos grandes)
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// remover
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
