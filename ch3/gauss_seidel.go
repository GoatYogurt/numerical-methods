package ch3

import (
	"fmt"
	"ppt/utils"
)

func Gauss_seidel(n int, a [][]float64, b []float64, x0 []float64, TOL float64, N int) {
	x := make([]float64, n)
	for k:= 1; k <= N; k++ {
		for i := 0; i < n; i++ {
			sum1 := 0.0
			sum2 := 0.0

			for j := 0; j < i; j++ {
				sum1 += a[i][j] * x[j]
			} 

			for j := i + 1; j < n; j++ {
				sum2 += a[i][j] * x0[j]
			}

			x[i] = 1/a[i][i] * (b[i] - sum1 - sum2)
		}

		if utils.EuclideanDistance(x, x0) < TOL {
			fmt.Printf("Roots: %v. Iterations: %d.\n", x, k)
			return
		}

		for i := 0; i < n; i++ {
			x0[i] = x[i]
		}
	}

	fmt.Printf("Max iterations exceeded.\n")
}
