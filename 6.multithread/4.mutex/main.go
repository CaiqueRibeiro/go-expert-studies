package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
Mutex é usado para sincronizar o acesso a uma variável compartilhada por várias goroutines.
Tudo entre m.Lock() e m.Unlock() é executado por uma única goroutine. Quando esta termina a operação,
libera a variável compartilhada para poder ser acessada por outra goroutine.
*/
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
