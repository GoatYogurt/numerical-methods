package ch7

import "fmt"

func Euler(a, b float64, N int, alpha float64, f func(float64, float64) float64) ([]float64, []float64) {
	h := (b-a)/float64(N)
	t := a
	w := alpha

	fmt.Printf("t: %f, w: %f\n", t, w)

	for i := 1; i <= N; i++ {
		w += h*f(t, w)
		t = a + float64(i)*h
		fmt.Printf("t: %f, w: %f\n", t, w)
	}
	return nil, nil
}