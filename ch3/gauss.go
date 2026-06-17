package ch3

import (
	"fmt"
	"math"
)

func swapRows(mat [][]float64, i, j int) {
	mat[i], mat[j] = mat[j], mat[i]
}

func Gauss(N int, augMatrix [][]float64) {
	for i := 0; i <= N-2; i++ {
		hasUniqueRoot := false
		p := i

		for r := i; r <= N-1; r++ {
			if math.Abs(augMatrix[r][i]) > 1e-9 {
				p = r
				hasUniqueRoot = true
				break
			}
		}

		if !hasUniqueRoot {
			fmt.Println("No unique solution exists.")
			return
		}

		if p != i {
			swapRows(augMatrix, p, i)
		}

		for j := i + 1; j <= N-1; j++ {
			m := augMatrix[j][i] / augMatrix[i][i]

			for k := i; k <= N; k++ {
				augMatrix[j][k] -= augMatrix[i][k] * m
			}
		}
	}

	if math.Abs(augMatrix[N-1][N-1]) < 1e-9 {
		fmt.Println("No unique solution exists.")
		return
	}

	x := make([]float64, N)
	x[N-1] = augMatrix[N-1][N] / augMatrix[N-1][N-1]

	for i := N - 2; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j <= N-1; j++ {
			sum += augMatrix[i][j] * x[j]
		}
		x[i] = (augMatrix[i][N] - sum) / augMatrix[i][i]
	}

	fmt.Printf("Roots: %v\n", x)
}
