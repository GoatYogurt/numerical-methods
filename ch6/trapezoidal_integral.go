package ch6

import (
	"errors"
	"ppt/utils"
)

// TrapezoidalIntegration tính gần đúng tích phân của f(x) trên đoạn [a, b] bằng phương pháp hình thang
// N: Số khoảng chia (N càng lớn, sai số càng nhỏ, bắt buộc N > 0)
func TrapezoidalIntegration(a, b float64, N int, f utils.IntegrableFunc) (float64, error) {
	if N <= 0 {
		return 0, errors.New("số khoảng chia N phải lớn hơn 0")
	}

	h := (b - a) / float64(N)
	
	// Khởi tạo tổng ban đầu bằng f(a) + f(b) theo công thức
	sum := f(a) + f(b)

	// Vòng lặp tính tổng các điểm nút nằm giữa (được nhân hệ số 2)
	for i := 1; i < N; i++ {
		xi := a + float64(i)*h
		sum += 2.0 * f(xi)
	}

	// Nhân với h / 2 ở bước cuối cùng để tối ưu hiệu năng tính toán
	integral := (h / 2.0) * sum

	return integral, nil
}