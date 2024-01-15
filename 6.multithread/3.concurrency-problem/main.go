package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number int64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		time.Sleep(200 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o usuário numero: %d\n", number)))
	})

	http.ListenAndServe(":8000", nil)
}

// Para solicitar que o go verifique possiveis RCs, execute o comando: go run -race main.go
