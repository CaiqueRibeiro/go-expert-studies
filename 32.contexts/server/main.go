package main

import (
	"log"
	"net/http"
	"time"
)

// CONTEXTS NO SERVIDOR
// usado para verificar que, caso o um contexto tenha sido cancelado ou com timeout aconteca alguma operacao

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada com sucesso")
	select {
	case <-time.After(5 * time.Second):
		// imprime no terminal
		log.Println("Request processada com sucesso")
		// imprime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		// imprime no terminal
		log.Println("Request cancelada pelo cliente")
	}
}
