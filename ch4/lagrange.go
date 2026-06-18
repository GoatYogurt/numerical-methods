package ch4

// import "fmt"

func LinearLagrange(x, x0, y0, x1, y1 float64) float64 {
	if x0 == x1 {
		panic("x0 and x1 are the same. Cannot interpolate!")
	}

	return y0 + (x-x0)*(y1-y0)/(x1-x0)
}

func l(k int, x float64, xs []float64) float64 {
	res := 1.0
	for i, xi := range xs {
		if i != k {
			res *= (x-xi)/(xs[k] - xi)
		}
	}
	return res
}

func GeneralLagrange(x float64, xs []float64, ys []float64) float64 {
	sum := 0.0
	n := len(xs)
	for k := 0; k < n; k++ {
		sum += ys[k] * l(k, x, xs)
	}
	return sum
}