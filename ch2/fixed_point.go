package main

import (
	"fmt"
	"math"
)

func fixed_point(p0 float64, g func(float64) float64, TOL float64, N int) {
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

func main() {
	// g(x) functions from page 61
	g1 := func(x float64) float64 {
		return x - x*x*x - 4*x*x + 10
	}

	g2 := func(x float64) float64 {
		return math.Pow(10/x - 4*x, 0.5)
	}

	g3 := func(x float64) float64 {
		return 0.5 * math.Pow(10 - x*x*x, 0.5)
	}

	g4 := func(x float64) float64 {
		return math.Pow(10/(x+4), 0.5)
	}

	g5 := func (x float64) float64 {
		return x - (x*x*x + 4*x*x - 10)/(3*x*x+8*x)
	}

	fixed_point(1.5, g1, 1e-100, 30)
	fixed_point(1.5, g2, 1e-100, 30)
	fixed_point(1.5, g3, 1e-100, 30)
	fixed_point(1.5, g4, 1e-100, 30)
	fixed_point(1.5, g5, 1e-100, 30)
}