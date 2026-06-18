package ch3

import (
	"fmt"
	"ppt/utils"
)

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

		if utils.EuclideanDistance(x, x0) < TOL {
			fmt.Printf("Roots: %v. Iterations: %d.\n", x, k)
			return
		}

		for i := 0; i < n; i++ {
			x0[i] = x[i]
		}
	}

	fmt.Printf("Max iterations exceeded.")
}
