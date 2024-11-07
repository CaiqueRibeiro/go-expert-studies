package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type PedidoServiceImpl struct {
	db *sql.DB
}

func NewPedidoServiceImpl() *PedidoServiceImpl {
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

	return &PedidoServiceImpl{db: db}
}

func (s *PedidoServiceImpl) GetPedidos(clienteID int) []Pedido {
	// Defina a consulta SQL para buscar os pedidos
	query := `SELECT numero, cliente_id, valor, bloqueado FROM pedidos WHERE cliente_id=$1`

	// Execute a consulta
	rows, err := s.db.Query(query, clienteID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pedidos []Pedido
	for rows.Next() {
		var p Pedido
		if err := rows.Scan(&p.Numero, &p.Cliente_id, &p.Valor, &p.Bloqueado); err != nil {
			log.Fatal(err)
		}
		pedidos = append(pedidos, p)
	}

	// Verifique por erros durante a iteração
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return pedidos
}
