package ch6

import (
	"errors"
	"math"
)

// FivePointMidpoint tính đạo hàm bậc nhất tại điểm nút trung tâm x0 
// dựa trên 4 điểm nút lân cận cách đều nhau một khoảng h.
// h: Khoảng cách bước lưới giữa các nút liên tiếp (bắt buộc h > 0)
// fMinus2h, fMinusH: Giá trị hàm số tại f(x0 - 2h) và f(x0 - h)
// fPlusH, fPlus2h: Giá trị hàm số tại f(x0 + h) và f(x0 + 2h)
func FivePointMidpoint(h float64, fMinus2h, fMinusH, fPlusH, fPlus2h float64) (float64, error) {
	// Kiểm tra điều kiện biên an toàn để tránh lỗi chia cho 0
	if math.Abs(h) < 1e-15 {
		return 0, errors.New("bước lưới h không hợp lệ (gần bằng 0), nguy cơ bùng nổ sai số do nhiễu số")
	}

	// Công thức năm điểm trung tâm: 
	// f'(x0) = [f(x0 - 2h) - 8f(x0 - h) + 8f(x0 + h) - f(x0 + 2h)] / (12h)
	derivative := (fMinus2h - 8.0*fMinusH + 8.0*fPlusH - fPlus2h) / (12.0 * h)

	return derivative, nil
}