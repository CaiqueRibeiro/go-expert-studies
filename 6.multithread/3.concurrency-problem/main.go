package main

import (
	"fmt"
	"net/http"
	"time"
)

var number int64 = 0 // Uma variável global pode ser acessada por todas as goroutines existentes em nível maior que ela

func main() {
	// Ai criar uma função handler de HTTP, cada requisição será feita em uma nova goroutine
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++ // manusear uma variável global em uma goroutine pode causar problemas de concorrência, onde mais de uma goroutine tenta acessar a mesma variável ao mesmo tempo
		time.Sleep(200 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o usuário numero: %d\n", number))) // O valor tende a ser menor do que o real número de conexões pois a variável global está sendo acessada por mais de uma goroutine ao mesmo tempo
	})

	http.ListenAndServe(":8000", nil)
}

// Para solicitar que o go verifique possiveis RCs, execute o comando: go run -race main.go
