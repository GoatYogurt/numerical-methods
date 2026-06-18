package ch3

import (
	"fmt"
	"math"
)

func EuclideanDistance(x, x0 []float64) float64 {
	if len(x) != len(x0) {
		panic("Two sclices must have same lengths.")
	}

	sumOfSquares := 0.0

	for i := 0; i < len(x); i++ {
		diff := x[i] - x0[i]
		sumOfSquares += diff * diff
	}

	return math.Sqrt(sumOfSquares)
}

func Jacobi(n int, a [][]float64, b []float64, x0 []float64, TOL float64, N int) {
	x := make([]float64, n)
	for k:= 1; k <= N; k++ {
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += a[i][j] * x0[j]
				}
			}
			x[i] = 1/a[i][i] * (b[i] - sum)
		}

		if EuclideanDistance(x, x0) < TOL {
			fmt.Printf("Roots: %v. Iterations: %d.\n", x, k)
			return
		}

		for i := 0; i < n; i++ {
			x0[i] = x[i]
		}
	}

	fmt.Printf("Max iterations exceeded.")
}
