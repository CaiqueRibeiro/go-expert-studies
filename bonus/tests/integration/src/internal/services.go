package internal

import "time"

type Parametros struct {
	ValorLimiteUltimoPedido    float64
	ValorLimiteUltimoPagamento float64
}

type Cliente struct {
	Id              int
	DataCadastro    time.Time
	SituacaoCredito string
}

type Pedido struct {
	Numero     int
	Cliente_id int
	Valor      float64
	Bloqueado  bool
}

type Pagamento struct {
	Numero       int
	Vencimento   time.Time
	Valor        float64
	Liquidado    bool
	Cliente_id   int
	Inadimplente bool
}
