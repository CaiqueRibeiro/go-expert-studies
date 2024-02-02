package main

import (
	"io"
	"net/http"
)

func main() {
	// Faz uma requisição a algum endpoint HTTP
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	// Lê o corpo da requisição (HTML, JSON, XML, etc)
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	// Vem como array de bytes
	println(string(res))
	// Fecha o canal de IO
	req.Body.Close()
}
