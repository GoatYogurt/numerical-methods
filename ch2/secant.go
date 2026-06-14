package ch2

import (
	"fmt"
	"math"
)


func Secant(p0 float64, p1 float64, f func(float64) float64, TOL float64, N int) {
	i := 2
	q0 := f(p0)
	q1 := f(p1)
	p := 0.0

	for i <= N {
		p = p1-q1*(p1-p0)/(q1-q0)
		if math.Abs(p - p1) < TOL {
			fmt.Printf("Succeeded. Root: %f. Iterations: %d.\n", p, i)
			return
		}
		i += 1
		p0 = p1
		q0 = q1
		p1 = p
		q1 = f(p)
	}

	fmt.Printf("Failed. Root: %f. Iterations: %d.\n", p, i)
}
