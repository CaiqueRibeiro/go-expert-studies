package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type PagamentoServiceImpl struct {
	db *sql.DB
}

func NewPagamentoServiceImpl() *PagamentoServiceImpl {
	// Obtenha as credenciais do banco de dados de variáveis de ambiente
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Conecte-se ao banco de dados
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Verifique a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &PagamentoServiceImpl{db: db}
}

func (s *PagamentoServiceImpl) GetPagamentosCliente(clienteID int) []Pagamento {
	// Defina a consulta SQL para buscar os pagamentos de um cliente
	query := `SELECT numero, vencimento, valor, liquidado, cliente_id, inadimplente FROM pagamentos WHERE cliente_id=$1`

	// Execute a consulta
	rows, err := s.db.Query(query, clienteID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pagamentos []Pagamento
	for rows.Next() {
		var p Pagamento
		if err := rows.Scan(&p.Numero, &p.Vencimento, &p.Valor, &p.Liquidado, &p.Cliente_id, &p.Inadimplente); err != nil {
			log.Fatal(err)
		}
		pagamentos = append(pagamentos, p)
	}

	// Verifique por erros durante a iteração
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return pagamentos
}

func main() {
	// Exemplo de como usar a PagamentoServiceImpl
	service := NewPagamentoServiceImpl()
	pagamentos := service.GetPagamentosCliente(1) // Assumindo 1 como um exemplo de clienteID
	fmt.Println(pagamentos)
}
