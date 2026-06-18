package main

import (
	"fmt"
	"ppt/ch4"
)

func main() {
	fmt.Println(ch4.Newton(
        []float64{1.0, 1.3, 1.6, 1.9, 2.2},
        []float64{0.7651977, 0.6200860, 0.4554022, 0.2818186, 0.1103623},
    ))


	// f := func (x float64) float64 {
	// 	return 1/x
	// }

	// fmt.Println(ch4.LinearLagrange(5,2,4,5,1))
	// fmt.Println(ch4.GeneralLagrange(
	// 	3,
	// 	[]float64 {2, 2.75, 4},
	// 	[]float64{f(2), f(2.75), f(4)},
	// 	),
	// )

	// ch3.Gauss_seidel(
	// 	4, 
	// 	[][]float64{
	// 		{10,-1,2,0},
	// 		{-1,11,-1,3},
	// 		{2,-1,10,-1},
	// 		{0,3,-1,8},
	// 	},
	// 	[]float64{6,25,-11,15},
	// 	[]float64{0,0,0,0},
	// 	1e-3,
	// 	100,
	// )

	// ch3.Jacobi(
	// 	4, 
	// 	[][]float64{
	// 		{10,-1,2,0},
	// 		{-1,11,-1,3},
	// 		{2,-1,10,-1},
	// 		{0,3,-1,8},
	// 	},
	// 	[]float64{6,25,-11,15},
	// 	[]float64{0,0,0,0},
	// 	1e-3,
	// 	11,
	// )
	
	// augMatrix := [][]float64{
	// 	{8,2,-7},
	// 	{3, 5, 2, 8},
	// 	{6,2,8,26},
	// }

	// ch3.Gauss(3, augMatrix)
	
	// TOL := 1e-8
	// N0 := 20
	// f := func(x float64) float64 {
	// 	return math.Cos(x) - x
	// }

	// ch2.False_position(0.5, math.Pi/4, f, TOL, N0)
	// ch2.Secant(0.5, math.Pi/4, f, TOL, N0)
	// ch2.Newton(math.Pi/4, f, TOL, N0)

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