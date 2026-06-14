package main

import (
	"fmt"
	"math"
)

func bisection(start float64, end float64, f func(float64) float64, TOL float64, N int) float64 {
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

// func main() {
// 	f := func(x float64) float64 {
// 		return math.Pow(x, 3) + 4*math.Pow(x, 2) - 10
// 	}
// 	fmt.Println(bisection(1, 2, f, math.Pow(10, -4), 100))
// }