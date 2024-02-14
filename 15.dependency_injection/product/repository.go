package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetProduct(id string) (Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProduct(id string) (Product, error) {
	return Product{
		Id:   id,
		Name: "Product One",
	}, nil
}
