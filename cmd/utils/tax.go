package utils

import (
	"runtime"
)

func CalculateTax(amount float64) float64 {

	runtime.GOMAXPROCS(1)

	if amount >= 1_000 {
		return 10.0
	}

	return 5.0
}
