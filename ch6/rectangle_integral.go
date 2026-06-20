package ch6

import (
	"errors"
	// "math"
	"ppt/utils"
)

// MidpointRectangleIntegration tính gần đúng tích phân của f(x) trên đoạn [a, b]
// N: Số khoảng chia (N càng lớn, sai số càng nhỏ, bắt buộc N > 0)
func MidpointRectangleIntegration(a, b float64, N int, f utils.IntegrableFunc) (float64, error) {
	if N <= 0 {
		return 0, errors.New("số khoảng chia N phải lớn hơn 0")
	}

	h := (b - a) / float64(N)
	sum := 0.0

	// Vòng lặp tính tổng diện tích các hình chữ nhật
	for i := 0; i < N; i++ {
		// Tính trung điểm của đoạn con thứ i
		xMid := a + (float64(i)+0.5)*h
		
		// Cộng dồn diện tích: Chiều cao f(xMid) * Chiều rộng h
		sum += f(xMid)
	}

	// Đưa h ra ngoài nhân một lần để tối ưu phép nhân số thực
	integral := sum * h

	return integral, nil
}