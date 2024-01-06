package main

// QUANDO USAR PONTEIROS
// Usar quando quiser que as variáveis seja mutáveis ao longo do código
// Usado também para economizar memória: ao invés de passar uma cópia de uma variável, passa-se o endereço dela

func soma(a, b int) int {
	// é feita uma CÓPIA dos parametros, e não mexe no original
	return a + b
}

func soma2(a, b *int) int {
	// usando o ponteiro como parâmetro, usa-se os valores reais e não mais uma cópia
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	soma(minhaVar1, minhaVar2)
	soma2(&minhaVar1, &minhaVar2)
}
