package main

import (
	"fmt"
	"net/http"
	"time"
)

var number int64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(200 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o usuário numero: %d\n", number)))
	})

	http.ListenAndServe(":8000", nil)
}
