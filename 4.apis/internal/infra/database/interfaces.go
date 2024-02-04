package database

import "github.com/CaiqueRibeiro/4-api/internal/entity"

type UserInterface interface {
	Create(u *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(p *entity.Product) error
	FindByID(id string) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	Update(p *entity.Product) error
	Delete(id string) error
}
