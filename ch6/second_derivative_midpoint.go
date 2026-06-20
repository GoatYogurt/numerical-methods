package ch6

import (
	"errors"
	"math"
)

// SecondDerivMidpoint tính đạo hàm bậc hai tại điểm nút trung tâm x0.
// h: Khoảng cách bước lưới giữa các điểm nút (bắt buộc h > 0)
// fMinusH: Giá trị hàm số tại điểm nút phía sau f(x0 - h)
// f0: Giá trị hàm số tại chính điểm nút trung tâm f(x0)
// fPlusH: Giá trị hàm số tại điểm nút phía trước f(x0 + h)
func SecondDerivMidpoint(h float64, fMinusH, f0, fPlusH float64) (float64, error) {
	// Kiểm tra điều kiện mẫu số để tránh lỗi chia cho 0
	if math.Abs(h) < 1e-15 {
		return 0, errors.New("bước lưới h không hợp lệ (gần bằng 0), mẫu số h^2 sẽ bị triệt tiêu")
	}

	// Công thức đạo hàm bậc hai điểm giữa:
	// f''(x0) = [f(x0 - h) - 2f(x0) + f(x0 + h)] / h^2
	secondDerivative := (fMinusH - 2.0*f0 + fPlusH) / (h * h)

	return secondDerivative, nil
}