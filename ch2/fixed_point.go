package ch2

import (
	"fmt"
	"math"
)

func Fixed_point(p0 float64, g func(float64) float64, TOL float64, N int) {
	i := 1
	p := 0.0

	for i <= N {
		p = g(p0)
		if math.Abs(p - p0) < TOL {
			fmt.Printf("Root: %f. Iterations: %d.\n", p, i)
			return
		}
		i += 1
		p0 = p
	}

	fmt.Printf("Root: %f. Iterations: %d.\n", p, i)
	return
}