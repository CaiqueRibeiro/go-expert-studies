package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

/*
	A biblioteca atomic fornece operações atômicas de baixo nível. Essas operações são úteis para implementar algoritmos de sincronização,
	permitindo que você evite o uso de mutexes.
*/

var number int64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&number, 1)
		time.Sleep(200 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o usuário numero: %d\n", number)))
	})

	http.ListenAndServe(":8000", nil)
}
