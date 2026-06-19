package ch7

import "ppt/utils"

// HeunSolver giải bài toán IVP bằng phương pháp Heun (RK2)
// HeunThirdOrderSolver giải phương trình vi phân bằng phương pháp Heun bậc 3
// a, b: Khoảng giải bài toán [a, b]
// N: Số lượng bước nhảy
// alpha: Điều kiện ban đầu w0 = alpha
// f: Hàm dy/dt = f(t, w)
func Heun(a, b float64, N int, alpha float64, f func(float64, float64) float64) []utils.Point {
	h := (b - a) / float64(N)
	points := make([]utils.Point, N+1)

	// Khởi tạo điểm đầu tiên: w0 = alpha (Theo Step 1)
	t := a
	w := alpha
	points[0] = utils.Point{T: t, W: w}

	// Vòng lặp tính toán cho i = 0, 1, ..., N-1
	for i := 0; i < N; i++ {
		// 1. Tính k1 từ trạng thái hiện tại (ti, wi)
		k1 := f(t, w)

		// 2. Tính k2 tại điểm phụ thứ nhất (ti + h/3)
		t2 := t + h/3.0
		w2 := w + (h/3.0)*k1
		k2 := f(t2, w2)

		// 3. Tính k3 tại điểm phụ thứ hai (ti + 2h/3)
		t3 := t + (2.0*h)/3.0
		w3 := w + ((2.0*h)/3.0)*k2
		k3 := f(t3, w3)

		// 4. Cập nhật giá trị nghiệm chính thức tiếp theo: w_(i+1)
		w = w + (h/4.0)*(k1+3.0*k3)
		
		// Cập nhật giá trị thời gian: t_(i+1)
		t = a + float64(i+1)*h

		// Lưu kết quả vào mảng
		points[i+1] = utils.Point{T: t, W: w}
	}

	return points
}