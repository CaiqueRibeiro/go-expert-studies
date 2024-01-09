package tax

import "errors"

func CalculateTax(price float64) (float64, error) {
	if price <= 0 {
		return 0.0, errors.New("amount must be greater than 0")
	}
	if price >= 1000 {
		return 10.0, nil
	} else {
		return 5.0, nil
	}
}