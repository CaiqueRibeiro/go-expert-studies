package main
// CRIAÇÃO DE TIPOS PRÓPRIOS

type ID int

var (
	c int // infere 0
	d string // infere ''
	e float64 // infere 0.00000
	f int = 10 // também é possível atribuir na declaração
	g ID
)

func main() {
	println(g)
}