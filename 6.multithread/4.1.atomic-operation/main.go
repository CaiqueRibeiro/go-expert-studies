package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number int64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&number, 1)
		time.Sleep(200 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o usuário numero: %d\n", number)))
	})

	http.ListenAndServe(":8000", nil)
}

// Para solicitar que o go verifique possiveis RCs, execute o comando: go run -race main.go
