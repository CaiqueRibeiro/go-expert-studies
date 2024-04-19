package repository

import (
	"context"
	"database/sql"

	"github.com/CaiqueRibeiro/unity-of-work/internal/db"
	"github.com/CaiqueRibeiro/unity-of-work/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(dtb *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}
