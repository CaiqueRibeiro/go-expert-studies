package main

import "fmt"

// INICIANDO COM STRUCTS

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	caique := Cliente{
		Nome:  "Caique",
		Idade: 28,
		Ativo: false,
	}

	caique.Ativo = true

	fmt.Printf("Nome:%s\nIdade: %d\nAtivo:%t", caique.Nome, caique.Idade, caique.Ativo)
}
