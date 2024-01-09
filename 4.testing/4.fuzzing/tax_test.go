package tax

import "testing"

/* Fuzzing é usado para gerar automaticamente vários conjuntos de números,
a fim de testar muitos cenários sem que seja preciso escrevê-los manualmente
*/
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0, 3940.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
		if(amount >= 1000 && result != 10.0) {
			t.Errorf("Received %f but expected 10.0", result)
		}
		if(amount < 1000 && amount > 0 && result != 5.0) {
			t.Errorf("Received %f but expected 5.0", result)
		}
	})
}