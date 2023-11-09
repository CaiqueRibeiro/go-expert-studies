package main

import "fmt"

// PONTEIROS E STRUCTS

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	// c é apenas uma cópia da struct que o chamou, e não a própria
	c.nome = "Caique Ribeiro" // usando (c Cliente) esta linha não vai alterar realmente a struct q o chamou
	fmt.Printf("\nO cliente %v andou", c.nome)
}

func (c *Cliente) andou_de_verdade() {
	c.nome = "Caique Ribeiro" // agora sim irá alterar a struct verdadeira que o chamou
	fmt.Printf("\nO cliente %v andou de verdade", c.nome)
}

func main() {
	caique := Cliente{
		nome: "Caique",
	}
	caique.andou()
	fmt.Printf("\nO valor da struct com nome é %s", caique.nome) // Continuará somente Caique

	caique.andou_de_verdade()
	fmt.Printf("\nO valor da struct com nome é %s", caique.nome) // Agora será Caique
}
