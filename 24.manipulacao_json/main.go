package main

import (
	"encoding/json"
	"os"
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
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta) // struct to json
	if err != nil {
		println(err)
	}
	println(string(res)) // JSON sempre retorna em bytes

	// Encoder: já transforma diretamente em JSON e retorna para algo (stdout, arquivo, webserver)
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(conta)
	if err != nil {
		println(err)
	}

	// Decodificar: transformar o JSON em uma struct ou variável
	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)

	// Para encode/decode quando structs e JSON não são iguais, usa-se as tags
	jsonPuro = []byte(`{"n":2,"s":200}`)
	var contaY ContaComTag
	err = json.Unmarshal(jsonPuro, &contaY)
	if err != nil {
		println(err)
	}
	println(contaY.Saldo)
}
