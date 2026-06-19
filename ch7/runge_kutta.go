package ch7

import (
	"fmt"
)

// RungeKutta4 triển khai thuật toán RK4 giải bài toán Cauchy (IVP)
// a, b: khoảng tính toán [a, b]
// N: số bước chia
// alpha: điều kiện ban đầu y(a) = alpha
// f: hàm vế phải f(t, y) đại diện cho đạo hàm dy/dt
func RungeKutta4(a, b float64, N int, alpha float64, f func(float64, float64) float64) {
	// Step 1: Thiết lập các giá trị ban đầu
	h := (b - a) / float64(N)
	t := a
	w := alpha

	// In ra giá trị ban đầu (t_0, w_0)
	fmt.Printf("Step 0: t = %.4f, w = %.6f\n", t, w)

	// Step 2: Vòng lặp cho i = 1, 2, ..., N ứng với Steps 3-5
	for i := 1; i <= N; i++ {
		// Step 3: Tính toán các hệ số K1, K2, K3, K4
		k1 := h * f(t, w)
		k2 := h * f(t+h/2.0, w+k1/2.0)
		k3 := h * f(t+h/2.0, w+k2/2.0)
		k4 := h * f(t+h, w+k3)

		// Step 4: Cập nhật w_i và t_i cho bước tiếp theo
		w = w + (k1+2.0*k2+2.0*k3+k4)/6.0
		t = a + float64(i)*h

		// Step 5: OUTPUT (t_i, w_i) của từng bước lặp
		fmt.Printf("Step %d: t = %.4f, w = %.6f\n", i, t, w)
	}
	// Step 6: STOP
}