package handlers

// Handlers são equivalentes aos controllers na arquitetura hexagonal. Eles são responsáveis por receber as requisições HTTP, chamar os casos de uso e retornar as respostas.
import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueRibeiro/4-api/internal/dto"
	"github.com/CaiqueRibeiro/4-api/internal/entity"
	"github.com/CaiqueRibeiro/4-api/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput              // usando um DTO para receber os dados do client
	err := json.NewDecoder(r.Body).Decode(&product) // decodifiando o JSON e já colocando em uma struct
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // caso dê erro, retorna um HTTP Bad Request pro cliente
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price) // Cria a entidade de produto
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p) // Salva no banco
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
