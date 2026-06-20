package ch6

import (
	"errors"
	"math"
)

// ThreePointEndpoint tính đạo hàm bậc nhất tại điểm nút đầu tiên x0 
// dựa trên giá trị của 3 điểm nút cách đều nhau một khoảng h.
// h: Khoảng cách giữa các bước lưới (bắt buộc h > 0 hoặc h < 0)
// f0, f1, f2: Giá trị của hàm số tại f(x0), f(x0+h), f(x0+2h)
func ThreePointEndpoint(h float64, f0, f1, f2 float64) (float64, error) {
	// Kiểm tra tránh lỗi chia cho 0 nếu bước lưới quá nhỏ hoặc bằng 0
	if math.Abs(h) < 1e-15 {
		return 0, errors.New("bước lưới h không hợp lệ (gần bằng 0), nguy cơ bùng nổ sai số nhiễu")
	}

	// Công thức ba điểm cực trị: f'(x0) = [-3f(x0) + 4f(x0+h) - f(x0+2h)] / (2h)
	derivative := (-3.0*f0 + 4.0*f1 - f2) / (2.0 * h)
	
	return derivative, nil
}