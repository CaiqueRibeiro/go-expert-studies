package internal_test

import (
	"testing"
	"time"

	internal "cleutonsampaio.com/golangdemo1/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock para o serviço de Pedido
type MockPedidoService struct {
	mock.Mock
}

func (m *MockPedidoService) GetPedidos(clienteID int) []internal.Pedido {
	args := m.Called(clienteID)
	return args.Get(0).([]internal.Pedido)
}

// Mock para o serviço de Pagamento
type MockPagamentoService struct {
	mock.Mock
}

func (m *MockPagamentoService) GetPagamentosCliente(clienteID int) []internal.Pagamento {
	args := m.Called(clienteID)
	return args.Get(0).([]internal.Pagamento)
}

// Preconditions test: VerificarFormasPagamento
func TestVerificarFormasPagamentoPreconditions(t *testing.T) {
	// Configurando mocks
	mockPedidoService := &MockPedidoService{}
	mockPagamentoService := &MockPagamentoService{}

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	cliente := internal.Cliente{Id: 1, DataCadastro: time.Now(), SituacaoCredito: "boa"}

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(mockPedidoService, mockPagamentoService, parametros)

	// Testando valorEntrada negativo
	assert.PanicsWithValue(t, "verificarFormasPagamento: Valor entrada inválido", func() {
		formasPagamentoService.VerificarFormasPagamento(cliente, -1, 1000)
	})

	// Testando valorEntrada maior que o valor do pedido
	assert.PanicsWithValue(t, "verificarFormasPagamento: Valor entrada maior que o valor do pedido", func() {
		formasPagamentoService.VerificarFormasPagamento(cliente, 2000, 1000)
	})

	// Testando valorPedido negativo
	assert.PanicsWithValue(t, "verificarFormasPagamento: Valor do pedido inválido", func() {
		formasPagamentoService.VerificarFormasPagamento(cliente, 0, -1)
	})
}

func TestVerificarFormasPagamento_cliente_menos_6_meses(t *testing.T) {
	// Configurando mocks
	mockPedidoService := &MockPedidoService{}
	mockPagamentoService := &MockPagamentoService{}

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	cliente := internal.Cliente{Id: 1, DataCadastro: time.Now(), SituacaoCredito: "boa"}

	// Configurando retorno dos mocks
	mockPedidoService.On("GetPedidos", 1).Return([]internal.Pedido{{Numero: 1, Valor: 100}, {Numero: 2, Valor: 500}})
	mockPagamentoService.On("GetPagamentosCliente", 1).Return([]internal.Pagamento{{Numero: 1, Valor: 300, Liquidado: true},
		{Numero: 2, Valor: 300, Liquidado: true}})

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(mockPedidoService, mockPagamentoService, parametros)

	// Chamando o método sob teste
	formasPagamento := formasPagamentoService.VerificarFormasPagamento(cliente, 100, 1000)

	// Verificando resultados
	assert.ElementsMatch(t, []string{"Pagamento à vista"}, formasPagamento)

	// Verificando chamadas dos mocks
	mockPedidoService.AssertCalled(t, "GetPedidos", cliente.Id)
	mockPagamentoService.AssertCalled(t, "GetPagamentosCliente", cliente.Id)
}

func TestVerificarFormasPagamento_cliente_credito_ruim(t *testing.T) {
	// Configurando mocks
	mockPedidoService := &MockPedidoService{}
	mockPagamentoService := &MockPagamentoService{}

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	// Obtém a data atual
	hoje := time.Now()
	// Calcula a data um ano atrás
	umAnoAtras := hoje.AddDate(-1, 0, 0)
	cliente := internal.Cliente{Id: 1, DataCadastro: umAnoAtras, SituacaoCredito: "ruim"}

	// Configurando retorno dos mocks
	mockPedidoService.On("GetPedidos", 1).Return([]internal.Pedido{{Numero: 1, Valor: 500}})
	mockPagamentoService.On("GetPagamentosCliente", 1).Return([]internal.Pagamento{{Numero: 1, Valor: 300, Liquidado: true}})

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(mockPedidoService, mockPagamentoService, parametros)

	// Chamando o método sob teste
	formasPagamento := formasPagamentoService.VerificarFormasPagamento(cliente, 100, 1000)

	// Verificando resultados
	assert.ElementsMatch(t, []string{"Pagamento à vista"}, formasPagamento)

	// Verificando chamadas dos mocks
	mockPedidoService.AssertCalled(t, "GetPedidos", cliente.Id)
	mockPagamentoService.AssertCalled(t, "GetPagamentosCliente", cliente.Id)
}

func TestVerificarFormasPagamento_cliente_credito_regular_com_6_meses_entrada_menor(t *testing.T) {
	// Configurando mocks
	mockPedidoService := &MockPedidoService{}
	mockPagamentoService := &MockPagamentoService{}

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	// Obtém a data atual
	hoje := time.Now()
	// Calcula a data um ano atrás
	seisMeses := hoje.AddDate(0, -6, 0)
	cliente := internal.Cliente{Id: 1, DataCadastro: seisMeses, SituacaoCredito: "regular"}

	// Configurando retorno dos mocks
	mockPedidoService.On("GetPedidos", 1).Return([]internal.Pedido{{Numero: 1, Valor: 500}})
	mockPagamentoService.On("GetPagamentosCliente", 1).Return([]internal.Pagamento{{Numero: 1, Valor: 300, Liquidado: true}})

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(mockPedidoService, mockPagamentoService, parametros)

	// Chamando o método sob teste
	formasPagamento := formasPagamentoService.VerificarFormasPagamento(cliente, 100, 1000)

	// Verificando resultados
	assert.ElementsMatch(t, []string{"Pagamento à vista"}, formasPagamento)

	// Verificando chamadas dos mocks
	mockPedidoService.AssertCalled(t, "GetPedidos", cliente.Id)
	mockPagamentoService.AssertCalled(t, "GetPagamentosCliente", cliente.Id)
}

func TestVerificarFormasPagamento_cliente_credito_regular_com_6_meses_entrada_ok(t *testing.T) {
	// Configurando mocks
	mockPedidoService := &MockPedidoService{}
	mockPagamentoService := &MockPagamentoService{}

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	// Obtém a data atual
	hoje := time.Now()
	// Calcula a data um ano atrás
	seisMeses := hoje.AddDate(0, -6, 0)
	cliente := internal.Cliente{Id: 1, DataCadastro: seisMeses, SituacaoCredito: "regular"}

	// Configurando retorno dos mocks
	mockPedidoService.On("GetPedidos", 1).Return([]internal.Pedido{{Numero: 1, Valor: 500}})
	mockPagamentoService.On("GetPagamentosCliente", 1).Return([]internal.Pagamento{{Numero: 1, Valor: 300, Liquidado: true}})

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(mockPedidoService, mockPagamentoService, parametros)

	// Chamando o método sob teste
	formasPagamento := formasPagamentoService.VerificarFormasPagamento(cliente, 200, 1000)

	// Verificando resultados
	assert.ElementsMatch(t, []string{"Pagamento à vista", "Pagamento em 2 vezes com juros"}, formasPagamento)

	// Verificando chamadas dos mocks
	mockPedidoService.AssertCalled(t, "GetPedidos", cliente.Id)
	mockPagamentoService.AssertCalled(t, "GetPagamentosCliente", cliente.Id)
}

func TestVerificarFormasPagamento_cliente_credito_bom_com_6_meses(t *testing.T) {
	// Configurando mocks
	mockPedidoService := &MockPedidoService{}
	mockPagamentoService := &MockPagamentoService{}

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	// Obtém a data atual
	hoje := time.Now()
	// Calcula a data um ano atrás
	seisMeses := hoje.AddDate(0, -6, 0)
	cliente := internal.Cliente{Id: 1, DataCadastro: seisMeses, SituacaoCredito: "boa"}

	// Configurando retorno dos mocks
	mockPedidoService.On("GetPedidos", 1).Return([]internal.Pedido{{Numero: 1, Valor: 200}, {Numero: 2, Valor: 1000}})
	mockPagamentoService.On("GetPagamentosCliente", 1).Return([]internal.Pagamento{{Numero: 1, Valor: 100, Liquidado: true},
		{Numero: 2, Valor: 500, Liquidado: true}})

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(mockPedidoService, mockPagamentoService, parametros)

	// Chamando o método sob teste
	formasPagamento := formasPagamentoService.VerificarFormasPagamento(cliente, 0, 1000)

	// Verificando resultados
	assert.ElementsMatch(t, []string{"Pagamento à vista", "Pagamento em 2 vezes com juros"}, formasPagamento)

	// Verificando chamadas dos mocks
	mockPedidoService.AssertCalled(t, "GetPedidos", cliente.Id)
	mockPagamentoService.AssertCalled(t, "GetPagamentosCliente", cliente.Id)
}

func TestVerificarFormasPagamento_cliente_credito_bom_com_12_meses(t *testing.T) {
	// Configurando mocks
	mockPedidoService := &MockPedidoService{}
	mockPagamentoService := &MockPagamentoService{}

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	// Obtém a data atual
	hoje := time.Now()
	// Calcula a data um ano e 11 meses atrás
	umAnoAtras := hoje.AddDate(0, 0, +1)
	umAnoAtras = umAnoAtras.AddDate(-2, 0, 0)
	cliente := internal.Cliente{Id: 1, DataCadastro: umAnoAtras, SituacaoCredito: "boa"}

	// Configurando retorno dos mocks
	mockPedidoService.On("GetPedidos", 1).Return([]internal.Pedido{{Numero: 1, Valor: 200}, {Numero: 2, Valor: 200}})
	mockPagamentoService.On("GetPagamentosCliente", 1).Return([]internal.Pagamento{{Numero: 1, Valor: 100, Liquidado: true},
		{Numero: 2, Valor: 100, Liquidado: true}})

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(mockPedidoService, mockPagamentoService, parametros)

	// Chamando o método sob teste
	formasPagamento := formasPagamentoService.VerificarFormasPagamento(cliente, 0, 1000)

	// Verificando resultados
	assert.ElementsMatch(t, []string{"Pagamento à vista", "Pagamento em 2 vezes com juros", "Pagamento em 3 vezes sem juros"}, formasPagamento)

	// Verificando chamadas dos mocks
	mockPedidoService.AssertCalled(t, "GetPedidos", cliente.Id)
	mockPagamentoService.AssertCalled(t, "GetPagamentosCliente", cliente.Id)
}

func TestVerificarFormasPagamento_cliente_credito_top(t *testing.T) {
	// Configurando mocks
	mockPedidoService := &MockPedidoService{}
	mockPagamentoService := &MockPagamentoService{}

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	// Obtém a data atual
	hoje := time.Now()
	// Calcula a data um ano atrás
	umAnoAtras := hoje.AddDate(-1, 0, 0)
	cliente := internal.Cliente{Id: 1, DataCadastro: umAnoAtras, SituacaoCredito: "boa"}

	// Configurando retorno dos mocks
	mockPedidoService.On("GetPedidos", 1).Return([]internal.Pedido{{Numero: 1, Valor: 100},
		{Numero: 2, Valor: 1000}})
	mockPagamentoService.On("GetPagamentosCliente", 1).Return([]internal.Pagamento{
		{Numero: 2, Valor: 501, Liquidado: true},
		{Numero: 1, Valor: 350, Liquidado: true}})

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(mockPedidoService, mockPagamentoService, parametros)

	// Chamando o método sob teste
	formasPagamento := formasPagamentoService.VerificarFormasPagamento(cliente, 0, 1000)

	// Verificando resultados
	assert.ElementsMatch(t, []string{"Pagamento à vista", "Pagamento em 2 vezes com juros", "Pagamento em 3 vezes sem juros", "Pagamento em 6 vezes sem juros"}, formasPagamento)

	// Verificando chamadas dos mocks
	mockPedidoService.AssertCalled(t, "GetPedidos", cliente.Id)
	mockPagamentoService.AssertCalled(t, "GetPagamentosCliente", cliente.Id)
}
