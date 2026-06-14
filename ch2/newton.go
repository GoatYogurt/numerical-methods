package ch2

import (
	"fmt"
	"math"
)

// calculate derivative of a function by central difference
func derivative(f func(x float64) float64) func(float64) float64 {
	h := 1e-7
	return func(x float64) float64 {
		return (f(x + h)-f(x - h))/(2*h)
	}
}

func Newton(p0 float64, f func(float64) float64, TOL float64, N int) {
	i := 1
	df := derivative(f)
	p := 0.0

	for i <= N {
		p = p0 - f(p0)/df(p0)
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

// func main() {

// }