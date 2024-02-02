package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "arquivo.txt"
	content := "Escrevendo conteúdo no arquivo gerado"
	// Criar arquivo e escrever algo
	CriarEscreverNoArquivo(&path, &content)
	// Ler o arquivo inteiro
	LerArquivoTodo(&path)
	// Ler o arquivo em stream (para arquivos grandes)
	LerArquivoEmStream(&path)
}

func CriarEscreverNoArquivo(path *string, content *string) {
	// Criar o arquivo no disco
	f, err := os.Create(*path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// escrever no arquivo criado
	tamanho, err := f.WriteString(*content) // pode ser apenas Write([]byte("Hello, World")) se for em bytes
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho)
}

func LerArquivoTodo(path *string) {
	// Esse método carrega o arquivo todo na memória. Pode causar problemas em arquivos grandes
	arquivo, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	// O arquivo vem como array de bytes. Transforme em string para entender o conteúdo
	fmt.Println((string(arquivo)))
}

func LerArquivoEmStream(path *string) {
	arquivo, err := os.Open(*path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 10) // define um array de bytes com 10 posições. Será a quantidade de caracteres por vez
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
