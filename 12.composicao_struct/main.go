package main

import "fmt"

// COMPOSIÇÃO DE STRUCTS

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco // Isso é composição. Mas posso criar uma propriedade do tipo endereco fazendo: Address Endereco
}

func main() {
	caique := Cliente{
		Nome:  "Caique",
		Idade: 28,
		Ativo: true,
	}

	caique.Cidade = "Poá" // OU caique.Endereco.Cidade
	caique.Estado = "SP"
	caique.Logradouro = "Rua 9 de Julho"
	caique.Numero = 79

	fmt.Printf("Nome:%s\nIdade: %d\nAtivo:%t", caique.Nome, caique.Idade, caique.Ativo)
}
