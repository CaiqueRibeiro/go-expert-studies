package main

// GENERICS

// Generic de forma mais básica
func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// CONSTRAINTS -> Mesma funcionalidade, mas usando constraint que encapsula os tipos
type MyNumber int
type Number interface {
	~int | ~float64 // colocar ~ permite que tipos personalizados correspondentes funcionem no Generics
}

func Soma2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Comparações fazendo Generics: Preciso usar comparable. Apenas compara a igualdade
// Para resolver isso é preciso usar constraints do pacote CONSTRAINTS
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Caique": 10000, "Juliana": 5000, "Jorge": 3900}
	m2 := map[string]float64{"Caique": 8000.20, "Juliana": 5000.80, "Jorge": 4300.50}

	m3 := map[string]MyNumber{"Caique": 10000, "Juliana": 5000, "Jorge": 3900}
	println(Soma(m))
	println(Soma(m2))
	println(Soma2(m3))
	println(Compara(10, 30)) // se for colocado int e float ao mesmo tempo irá dar exceptions
}
