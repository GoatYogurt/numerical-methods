package ch4

import "fmt"

func Newton(xs []float64, ys []float64) []float64 {
	n := len(xs) - 1
	F := make([][]float64, n+1)
	
	for i:= 0; i <= n; i++ {
		F[i] = make([]float64, n+1)
		F[i][0] = ys[i]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			F[i][j] = (F[i][j-1]-F[i-1][j-1])/(xs[i]-xs[i-j])
		}
	}

	res := make([]float64, n+1)
	for i:= 0; i <= n; i++ {
		res[i] = F[i][i]
	}

	fmt.Println(F)

	return res
}