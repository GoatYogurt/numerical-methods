package ch5

import (
	// "math"
	"ppt/utils"
)
// Simpson13 tính tích phân xấp xỉ của hàm g(x) trên đoạn [a, b]
func Simpson13(a, b float64, intervals int, g utils.TargetFunc) float64 {
	if intervals%2 != 0 {
		intervals++
	}
	h := (b - a) / float64(intervals)
	sum := g(a) + g(b)

	for i := 1; i < intervals; i++ {
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * g(x)
		} else {
			sum += 4 * g(x)
		}
	}
	return (h / 3.0) * sum
}

// LegendrePolynomial định nghĩa giá trị của đa thức Legendre tiêu chuẩn P_n(x) trên đoạn [-1, 1]
// Sử dụng công thức truy hồi Bonnet: P_n(x) = ((2n-1)xP_(n-1)(x) - (n-1)P_(n-2)(x)) / n
func LegendrePolynomial(n int, x float64) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	p0 := 1.0
	p1 := x
	var pn float64

	for i := 2; i <= n; i++ {
		floatI := float64(i)
		pn = ((2.0*floatI - 1.0) * x * p1 - (floatI - 1.0) * p0) / floatI
		p0 = p1
		p1 = pn
	}
	return p1
}

// OrthogonalLeastSquares tính toán bộ hệ số c_i của các đa thức trực giao dịch chuyển
// Trả về một mảng chứa các hệ số c0, c1, ..., cn
func OrthogonalLeastSquares(a, b float64, degree int, f utils.TargetFunc) []float64 {
	n := degree + 1
	c := make([]float64, n)
	const steps = 200 // Số khoảng chia tích phân Simpson

	// Hàm ánh xạ biến từ x thuộc [a, b] sang t thuộc [-1, 1] để dùng Legendre gốc
	// t = (2x - a - b) / (b - a)
	// dx = ((b - a) / 2) * dt
	
	for i := 0; i < n; i++ {
		currentDegree := i

		// Tích phân tử số: \int_{-1}^{1} f(x(t)) * P_i(t) dt
		numerator := Simpson13(-1.0, 1.0, steps, func(t float64) float64 {
			// Khôi phục x từ t: x = 0.5 * ((b - a) * t + a + b)
			x := 0.5 * ((b-a)*t + a + b)
			return f(x) * LegendrePolynomial(currentDegree, t)
		})

		// Tích phân mẫu số của đa thức Legendre tiêu chuẩn tính bằng công thức giải tích: 2 / (2i + 1)
		denominator := 2.0 / (2.0*float64(currentDegree) + 1.0)

		// Hệ số c_i
		c[i] = numerator / denominator
	}

	return c
}

// EvaluateOrthogonalPoly định giá hàm đa thức trực giao tại điểm x bất kỳ trên [a, b]
func EvaluateOrthogonalPoly(a, b float64, c []float64, x float64) float64 {
	// Chuyển x sang không gian t của Legendre [-1, 1]
	t := (2.0*x - a - b) / (b - a)
	sum := 0.0
	for i, coeff := range c {
		sum += coeff * LegendrePolynomial(i, t)
	}
	return sum
}