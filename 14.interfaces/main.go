package main

import "fmt"

// INTERFACES

// Qualquer struct que possuir os métodos da interface estará a "implementando"
type Pessoa interface {
	Desativar()
}

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
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	caique := Cliente{
		Nome:  "Caique",
		Idade: 28,
		Ativo: true,
	}

	Desativacao(caique) // posso passar pois Cliente implementa automaticamente a interface Pessoa
}
