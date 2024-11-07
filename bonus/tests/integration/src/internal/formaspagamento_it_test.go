//go:build integration
// +build integration

package internal_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"log"
	"os"
	"testing"
	"time"

	internal "cleutonsampaio.com/golangdemo1/internal"

	_ "github.com/lib/pq"

	postgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var postgresContainer *postgres.PostgresContainer
var ctx context.Context

func TestVerificarFormasPagamento_cliente_credito_bom_com_6_meses(t *testing.T) {
	// Configurando servicos
	pedidoService := internal.NewPedidoServiceImpl()
	pagamentoService := internal.NewPagamentoServiceImpl()

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	// Obtém a data atual
	hoje := time.Now()
	// Calcula a data 6 meses atrás
	meioAnoAtras := hoje.AddDate(0, -6, 0)
	cliente := internal.Cliente{Id: 100, DataCadastro: meioAnoAtras, SituacaoCredito: "boa"}

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(pedidoService, pagamentoService, parametros)

	// Chamando o método sob teste
	formasPagamento := formasPagamentoService.VerificarFormasPagamento(cliente, 0, 600)

	// Verificando resultados
	assert.ElementsMatch(t, []string{"Pagamento à vista", "Pagamento em 2 vezes com juros"}, formasPagamento)
}

func TestVerificarFormasPagamento_cliente_credito_top(t *testing.T) {
	// Configurando servicos
	pedidoService := internal.NewPedidoServiceImpl()
	pagamentoService := internal.NewPagamentoServiceImpl()

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	// Obtém a data atual
	hoje := time.Now()
	// Calcula a data um ano atrás
	umAnoAtras := hoje.AddDate(-1, 0, 0)
	cliente := internal.Cliente{Id: 200, DataCadastro: umAnoAtras, SituacaoCredito: "boa"}

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(pedidoService, pagamentoService, parametros)

	// Chamando o método sob teste
	formasPagamento := formasPagamentoService.VerificarFormasPagamento(cliente, 0, 1000)

	// Verificando resultados
	assert.ElementsMatch(t, []string{"Pagamento à vista", "Pagamento em 2 vezes com juros", "Pagamento em 3 vezes sem juros", "Pagamento em 6 vezes sem juros"}, formasPagamento)
}

func TestClienteComPedidoBloqueado(t *testing.T) {
	// Configurando servicos
	pedidoService := internal.NewPedidoServiceImpl()
	pagamentoService := internal.NewPagamentoServiceImpl()

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	cliente := internal.Cliente{Id: 300, DataCadastro: time.Now(), SituacaoCredito: "boa"}

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(pedidoService, pagamentoService, parametros)

	// Testando valorEntrada negativo
	assert.PanicsWithValue(t, "verificarFormasPagamento: Cliente com pedido bloqueado", func() {
		formasPagamentoService.VerificarFormasPagamento(cliente, 0, 1000)
	})
}

func TestClienteComPagamentoInadimplente(t *testing.T) {
	// Configurando servicos
	pedidoService := internal.NewPedidoServiceImpl()
	pagamentoService := internal.NewPagamentoServiceImpl()

	// Configurando parâmetros e cliente
	parametros := internal.Parametros{ValorLimiteUltimoPedido: 1000, ValorLimiteUltimoPagamento: 500}
	cliente := internal.Cliente{Id: 400, DataCadastro: time.Now(), SituacaoCredito: "boa"}

	// Inicializando o serviço
	formasPagamentoService := internal.NewFormasPagamentoService(pedidoService, pagamentoService, parametros)

	// Testando valorEntrada negativo
	assert.PanicsWithValue(t, "verificarFormasPagamento: Cliente com pagamento inadimplente", func() {
		formasPagamentoService.VerificarFormasPagamento(cliente, 0, 1000)
	})
}

func setup() {
	ctx := context.Background()
	dbUser := "postgres"
	dbPassword := "password"
	dbName := "postgres"
	os.Setenv("DB_USER", dbUser)
	os.Setenv("DB_PASSWORD", dbPassword)
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "postgres")
	var err error
	postgresContainer, err = postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:latest"),
		postgres.WithInitScripts("./initdb.sql"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(50*time.Second)),
	)
	hostport, _ := postgresContainer.MappedPort(ctx, "5432")
	os.Setenv("DB_PORT", hostport.Port())
	fmt.Println(hostport.Port())
	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}
}

func shutdown() {
	ctx := context.Background()
	if err := postgresContainer.Terminate(ctx); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
