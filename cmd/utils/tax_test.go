package utils_test

import (
	"testing"

	"github.com/IcaroSilvaFK/go_example_tests/cmd/utils"
)

func TestTax(t *testing.T) {

	amount := 1_000.0
	expect := 10.0

	r := utils.CalculateTax(amount)

	if r != expect {
		t.Error("Expected 10.0")
	}

}

func TestTaxWithAmountLessThan(t *testing.T) {

	amount := 500.0
	expect := 5.0

	r := utils.CalculateTax(amount)

	if r != expect {
		t.Error("Expected 5.0")
	}

}

func TestCalculateTaxBatch(t *testing.T) {
	type calTax struct {
		amount, expect float64
	}

	table := []calTax{
		{500.0, 5.0},
		{1_000.0, 10.0},
		{1_500.0, 10.0},
	}

	for _, v := range table {
		r := utils.CalculateTax(v.amount)

		if r != v.expect {
			t.Errorf("expected %v, but received %v", v.expect, r)
		}
	}

}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.CalculateTax(1_000.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{500.0, 1_000.0, 1_500.0, 1_000_000.0, 1_000_000_000.0}

	for _, i := range seed {
		f.Add(i)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		r := utils.CalculateTax(amount)
		if amount < 1_000.0 {
			if r != 5.0 {
				t.Errorf("expected 5.0, but received %v", r)
			}
		}
		if amount >= 1_000.0 {
			if r != 10.0 {
				t.Errorf("expected 10.0, but received %v", r)
			}
		}

	})
}
