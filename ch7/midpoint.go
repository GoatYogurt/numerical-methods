package ch7

import "ppt/utils"

// MidpointSolver giải bài toán IVP bằng phương pháp điểm giữa (RK2)
func Midpoint(a, b float64, N int, alpha float64, f func(float64, float64) float64) []utils.Point {
	h := (b - a) / float64(N)
	points := make([]utils.Point, N+1)

	t := a
	w := alpha
	points[0] = utils.Point{T: t, W: w}

	for i := 1; i <= N; i++ {
		k1 := f(t, w)
		wMid := w + (h/2.0)*k1
		tMid := t + (h / 2.0)

		k2 := f(tMid, wMid)
		w = w + h*k2
		t = a + float64(i)*h

		points[i] = utils.Point{T: t, W: w}
	}

	return points
}