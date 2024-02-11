package product

type ProductUseCase struct {
	repository *ProductRepository
}

func NewProductUseCase(r *ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository: r}
}

func (u *ProductUseCase) GetProduct(id string) (Product, error) {
	return u.repository.GetProduct(id)
}
