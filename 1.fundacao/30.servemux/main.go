package main

import (
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// as duas linhas a seguir atacham a rota no multiplexer (mux) padrao do Go. A vantagem eh a simplicidade e ele é global
	// http.HandleFunc("/", BuscaCepHandler)
	// http.ListenAndServe(":8080", nil)

	// Criando o proprio multiplexed (mux)
	// Ele nos obriga a ter um metodo ServeHTTP em uma struct e nos permite flexiblidade muito maior
	// Inclusive, no MUX padrao pode haver injecao de rotas de forma desenfreada
	mux := http.NewServeMux()
	mux.Handle("/blog", blog{title: "My blog"})
	http.ListenAndServe(":8081", mux)
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
