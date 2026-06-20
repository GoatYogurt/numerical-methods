package ch6

import (
	"errors"
	"math"
)

// FivePointEndpoint tính đạo hàm bậc nhất tại điểm nút biên x0
// dựa trên giá trị của 5 điểm nút cách đều nhau một khoảng h.
// h: Khoảng cách bước lưới (bắt buộc h > 0 hoặc h < 0)
// f0, f1, f2, f3, f4: Giá trị hàm số tại f(x0), f(x0+h), f(x0+2h), f(x0+3h), f(x0+4h)
func FivePointEndpoint(h float64, f0, f1, f2, f3, f4 float64) (float64, error) {
	// Kiểm tra điều kiện biên an toàn để tránh lỗi chia cho 0
	if math.Abs(h) < 1e-15 {
		return 0, errors.New("bước lưới h không hợp lệ (gần bằng 0), nguy cơ bùng nổ sai số do nhiễu số")
	}

	// Công thức năm điểm cực trị bậc O(h^4):
	// f'(x0) = [-25f0 + 48f1 - 36f2 + 16f3 - 3f4] / (12h)
	derivative := (-25.0*f0 + 48.0*f1 - 36.0*f2 + 16.0*f3 - 3.0*f4) / (12.0 * h)

	return derivative, nil
}