package ch2

import (
	"math"
)

func Bisection(start float64, end float64, f func(float64) float64, TOL float64, N int) float64 {
	i := 1
	fa := f(start)
	for i <= N {
		p := start+(end-start)/2
		fp := f(p)
		if f(p) == 0 || (end - start)/2 < TOL {
			return p
		}

		i += 1

		if fa*fp > 0 {
			start = p
			fa = fp
		} else {
			end = p
		}
	}
	return math.NaN()
}