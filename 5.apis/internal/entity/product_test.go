package entity

import (
	"testing"

	"github.com/CaiqueRibeiro/4-api/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	t.Run("It should create a new product", func(t *testing.T) {
		p, err := NewProduct("Product 1", 100)
		assert.Nil(t, err)
		assert.Equal(t, "Product 1", p.Name)
		assert.Equal(t, 100.00, p.Price)
		assert.NotEmpty(t, p.ID)
	})

	t.Run("It should return error when name is empty", func(t *testing.T) {
		p, err := NewProduct("", 100)
		assert.Equal(t, ErrNameIsRequired, err)
		assert.Nil(t, p)
	})

	t.Run("It should throw error when price is zero", func(t *testing.T) {
		p, err := NewProduct("Product 1", 0)
		assert.Equal(t, ErrPriceIsRequired, err)
		assert.Nil(t, p)
	})

	t.Run("It should throw error when price is invalid", func(t *testing.T) {
		p, err := NewProduct("Product 1", -1)
		assert.Equal(t, ErrPriceIsInvalid, err)
		assert.Nil(t, p)
	})

	t.Run("It should validate the product", func(t *testing.T) {
		p := &Product{
			ID:    entity.NewID(),
			Name:  "Product 1",
			Price: 100,
		}
		assert.Nil(t, p.Validate())
	})
}
