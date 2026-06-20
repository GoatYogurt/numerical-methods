package ch6

import (
	"errors"
	"math"
)

// ThreePointMidpoint tính đạo hàm bậc nhất tại điểm nút trung tâm x0.
// h: Khoảng cách bước lưới giữa các điểm nút (bắt buộc h > 0)
// fMinusH: Giá trị hàm số tại điểm nút phía sau f(x0 - h)
// fPlusH: Giá trị hàm số tại điểm nút phía trước f(x0 + h)
func ThreePointMidpoint(h float64, fMinusH, fPlusH float64) (float64, error) {
	// Kiểm tra tránh lỗi chia cho 0 nếu bước lưới gần bằng không
	if math.Abs(h) < 1e-15 {
		return 0, errors.New("bước lưới h không hợp lệ (gần bằng 0), nguy cơ bùng nổ sai số do nhiễu số")
	}

	// Công thức ba điểm trung tâm: f'(x0) = [f(x0 + h) - f(x0 - h)] / (2h)
	derivative := (fPlusH - fMinusH) / (2.0 * h)

	return derivative, nil
}