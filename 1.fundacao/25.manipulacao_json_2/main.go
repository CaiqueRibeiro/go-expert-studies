package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Conta struct {
	Numero int
	Saldo  int
}

type ContaComTag struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	// Faz uma requisição a algum endpoint HTTP
	req, err := http.Get("https://recupera_conta.com.br")
	if err != nil {
		panic(err)
	}
	// Lê o corpo da requisição
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var contaX Conta
	err = json.Unmarshal(res, &contaX) // transforma o JSON recebido em uma struct
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}
