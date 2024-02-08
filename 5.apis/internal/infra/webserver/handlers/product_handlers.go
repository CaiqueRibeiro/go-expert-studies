package handlers

// Handlers são equivalentes aos controllers na arquitetura hexagonal. Eles são responsáveis por receber as requisições HTTP, chamar os casos de uso e retornar as respostas.
import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CaiqueRibeiro/4-api/internal/dto"
	"github.com/CaiqueRibeiro/4-api/internal/entity"
	"github.com/CaiqueRibeiro/4-api/internal/infra/database"
	entityPkg "github.com/CaiqueRibeiro/4-api/pkg/entity"
	"github.com/go-chi/chi"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// Usando o roteador chi para as rotas
func (h *ProductHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")  // Pegando o parâmetro da URL
	pageInt, err := strconv.Atoi(page) // Convertendo para inteiro
	if err != nil {
		pageInt = 0 // Sem paginação
	}
	limit := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	order := r.URL.Query().Get("order")
	products, err := h.ProductDB.FindAll(pageInt, limitInt, order) // Buscando todos os produtos
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products) // Retornando os produtos em JSON
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // Pegando o id da URL
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id) // Buscando o produto no banco
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product) // Retornando o produto em JSON
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseID(id) // Verifica se o ID é um UUID
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
