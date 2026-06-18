package utils

import (
	"math"
)

func EuclideanDistance(x, x0 []float64) float64 {
	if len(x) != len(x0) {
		panic("Two slices must have same lengths.")
	}

	sumOfSquares := 0.0

	for i := 0; i < len(x); i++ {
		diff := x[i] - x0[i]
		sumOfSquares += diff * diff
	}

	return math.Sqrt(sumOfSquares)
}