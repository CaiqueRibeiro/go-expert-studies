package main

import "fmt"

// MÉTODOS EM STRUCTS

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

// O que estiver entre os parentes antes do nome delimitam a struct a qual pertence o método
func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {
	caique := Cliente{
		Nome:  "Caique",
		Idade: 28,
		Ativo: true,
	}

	caique.Address.Cidade = "Poá" // OU caique.Endereco.Cidade
	caique.Address.Estado = "SP"
	caique.Address.Logradouro = "Rua 9 de Julho"
	caique.Address.Numero = 79

	caique.Desativar()
}
