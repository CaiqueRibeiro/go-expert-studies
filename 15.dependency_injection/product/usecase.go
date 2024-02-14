package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(r ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository: r}
}

func (u *ProductUseCase) GetProduct(id string) (Product, error) {
	return u.repository.GetProduct(id)
}
