package tax

func CalculateTax(price float64) float64 {
	if price <= 0 {
		return 0.0
	}
	if price >= 1000 {
		return 10.0
	} else {
		return 5.0
	}
}

type Repository interface {
	SaveTax(tax float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax(amount)
	err := repository.SaveTax(tax)
	if err != nil {
		return err
	}
	return nil
}
