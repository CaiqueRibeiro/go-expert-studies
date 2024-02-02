package main

const a = "Hello, World!"

// sintaxe 1
var b bool // declaração sem atribuição (padrão false)

// sintaxe 2
var (
	c int          // infere 0
	d string       // infere ''
	e float64      // infere 0.00000
	f int     = 10 // também é possível atribuir na declaração
)

func main() {
	var a string = "X" // Variáveis declaradas em scopo LOCAL, caso não usadas, gerarão erro na compilação. Globais não levarão à isso.
	// Pode-se transformar a sintaxe acima em: a := "X"
	println(a)
}
