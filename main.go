package main

import (
	"math"

	"ppt/ch2"
)

func main() {
	TOL := 1e-8
	N0 := 20
	f := func(x float64) float64 {
		return math.Cos(x) - x
	}

	ch2.False_position(0.5, math.Pi/4, f, TOL, N0)
	ch2.Secant(0.5, math.Pi/4, f, TOL, N0)
	ch2.Newton(math.Pi/4, f, TOL, N0)

	// bisection
	// f := func(x float64) float64 {
	// 	return math.Pow(x, 3) + 4*math.Pow(x, 2) - 10
	// }
	// fmt.Println(ch2.Bisection(1, 2, f, math.Pow(10, -4), 100))

	// fixed point
	// // g(x) functions from page 61
	// g1 := func(x float64) float64 {
	// 	return x - x*x*x - 4*x*x + 10
	// }

	// g2 := func(x float64) float64 {
	// 	return math.Pow(10/x - 4*x, 0.5)
	// }

	// g3 := func(x float64) float64 {
	// 	return 0.5 * math.Pow(10 - x*x*x, 0.5)
	// }

	// g4 := func(x float64) float64 {
	// 	return math.Pow(10/(x+4), 0.5)
	// }

	// g5 := func (x float64) float64 {
	// 	return x - (x*x*x + 4*x*x - 10)/(3*x*x+8*x)
	// }

	// ch2.Fixed_point(1.5, g1, 1e-100, 30)
	// ch2.Fixed_point(1.5, g2, 1e-100, 30)
	// ch2.Fixed_point(1.5, g3, 1e-100, 30)
	// ch2.Fixed_point(1.5, g4, 1e-100, 30)
	// ch2.Fixed_point(1.5, g5, 1e-100, 30)
}