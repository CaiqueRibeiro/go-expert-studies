package main

import "fmt"

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

/*
Uma interface que define como as Optional Functions passadas para o Server
devem ser
*/
type OptFunc func(*Opts)

/*
Uma implementação da Optional Function (OptFunc) que habilita o TLS
*/
func withTLS(opts *Opts) {
	opts.tls = true
}

/*
Uma função factory que, usando o recurso de closure, returna uma função que
implementa a interface OptFunc e, por isso, pode ser passada como parâmetro
para o Server.

Ao usar functional option que precisem de parâmetro é preciso usar esta abordagem
e invocar a função ao invés de passá-la como parâmetro
*/
func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

type Server struct {
	Opts
}

/*
Se não for especificada nenhuma option, as padrões serão usadas.
É possivel passar infinitas funcional options, que irão usar o objeto Opts padrão
e incrementá-lo, já que ele é passado como parâmetro para cada uma delas como referência
*/
func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	s := newServer(withTLS, withMaxConn(99))
	fmt.Printf("%+v\n", s)
}
