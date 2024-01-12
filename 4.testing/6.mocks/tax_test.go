package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax := CalculateTax(500.0)
	assert.Equal(t, 5.0, tax)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil) // é possivel adicionar o .Once() no final para garantir que a função só será chamada uma vez
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	// repository.On("SaveTax", mock.Anything).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
	repository.AssertCalled(t, "SaveTax", 10.0)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
