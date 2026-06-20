package ch5

import (
	"errors"
	"math"
	"ppt/utils"
)

// Result đại diện cho hệ số của đường thẳng kết quả y = ax + b
type Result struct {
	A float64 // Hệ số góc (Slope)
	B float64 // Hệ số chặn (Intercept)
}

// LinearLeastSquares tìm đường thẳng y = ax + b tiệm cận gần nhất với tập dữ liệu rời rạc.
// X, Y: Các mảng tọa độ điểm đầu vào
func LinearLeastSquares(X []float64, Y []float64) (Result, error) {
	m := len(X)
	
	// Kiểm tra điều kiện ràng buộc dữ liệu đầu vào
	if m == 0 || len(Y) != m {
		return Result{}, errors.New("kích thước hai mảng dữ liệu X và Y không hợp lệ hoặc không bằng nhau")
	}
	if m < 2 {
		return Result{}, errors.New("cần tối thiểu 2 điểm dữ liệu để thực hiện xấp xỉ đường thẳng")
	}

	var sumX, sumY, sumX2, sumXY float64
	floatM := float64(m)

	// Vòng lặp O(m) tính toán các giá trị tổng tích lũy
	for i := 0; i < m; i++ {
		sumX += X[i]
		sumY += Y[i]
		sumX2 += X[i] * X[i]
		sumXY += X[i] * Y[i]
	}

	// Tính định thức mẫu số (Denominator) dựa trên quy tắc Cramer
	denominator := floatM*sumX2 - (sumX * sumX)

	// Kiểm tra tránh lỗi chia cho 0 (Khi tất cả các điểm X trùng nhau tạo thành đường thẳng đứng)
	if math.Abs(denominator) < 1e-15 {
		return Result{}, errors.New("mẫu số bằng 0 (các điểm X trùng nhau), hệ phương trình pháp tuyến vô nghiệm")
	}

	// Giải tìm hệ số a và b
	a := (floatM*sumXY - sumX*sumY) / denominator
	b := (sumX2*sumY - sumX*sumXY) / denominator

	return Result{A: a, B: b}, nil
}

// SimpsonIntegral tính xấp xỉ tích phân của hàm g(x) trên đoạn [a, b] bằng quy tắc Simpson 1/3
func SimpsonIntegral(a, b float64, intervals int, g utils.TargetFunc) float64 {
	if intervals%2 != 0 {
		intervals++ // Số khoảng chia phải là số chẵn
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

// GaussElimination giải hệ phương trình tuyến tính Ax = B với chiến lược phần tử trội hàng
func GaussElimination(A [][]float64, B []float64) ([]float64, error) {
	n := len(B)
	// Tạo ma trận mở rộng [A|B]
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n+1)
		for j := 0; j < n; j++ {
			matrix[i][j] = A[i][j]
		}
		matrix[i][n] = B[i]
	}

	// Tiến trình khử xuôi
	for i := 0; i < n; i++ {
		// Tìm phần tử trội trên cột i
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(matrix[k][i]) > math.Abs(matrix[maxRow][i]) {
				maxRow = k
			}
		}
		// Đổi chỗ hàng
		matrix[i], matrix[maxRow] = matrix[maxRow], matrix[i]

		if math.Abs(matrix[i][i]) < 1e-15 {
			return nil, errors.New("ma trận hệ số bị suy biến hoặc vô nghiệm")
		}

		// Khử các hàng phía dưới
		for k := i + 1; k < n; k++ {
			factor := matrix[k][i] / matrix[i][i]
			for j := i; j <= n; j++ {
				matrix[k][j] -= factor * matrix[i][j]
			}
		}
	}

	// Tiến trình thế ngược
	X := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += matrix[i][j] * X[j]
		}
		X[i] = (matrix[i][n] - sum) / matrix[i][i]
	}

	return X, nil
}

// ContinuousLeastSquares tìm các hệ số [a0, a1, ..., an] của đa thức xấp xỉ liên tục liên tục
// a, b: khoảng liên tục [a, b]
// degree: bậc n của đa thức cần tìm (ví dụ: degree = 1 là đường thẳng, 2 là parabol)
// f: hàm số gốc cần xấp xỉ
func ContinuousLeastSquares(a, b float64, degree int, f utils.TargetFunc) ([]float64, error) {
	n := degree + 1 // Số lượng ẩn số (a0 đến an)
	
	// Khởi tạo ma trận hệ số A và vector vế phải B
	A := make([][]float64, n)
	for i := range A {
		A[i] = make([]float64, n)
	}
	B := make([]float64, n)

	const integrationSteps = 200 // Số khoảng chia để tính tích phân Simpson cho mịn

	// 1. Tính toán các tích phân cho ma trận hệ số vế trái
	// Nhận xét: Tích phân của x^(i+j) chỉ phụ thuộc vào tổng lũy thừa (i+j)
	powers := make([]float64, 2*degree+1)
	for p := 0; p <= 2*degree; p++ {
		powerValue := p
		powers[p] = SimpsonIntegral(a, b, integrationSteps, func(x float64) float64 {
			return math.Pow(x, float64(powerValue))
		})
	}

	// Điền giá trị vào ma trận hàng-cột A
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A[i][j] = powers[i+j]
		}
	}

	// 2. Tính toán tích phân cho vector vế phải B: tích phân của x^i * f(x)
	for i := 0; i < n; i++ {
		currentPower := i
		B[i] = SimpsonIntegral(a, b, integrationSteps, func(x float64) float64 {
			return math.Pow(x, float64(currentPower)) * f(x)
		})
	}

	// 3. Giải hệ phương trình bằng khử Gauss để lấy các hệ số
	coefficients, err := GaussElimination(A, B)
	if err != nil {
		return nil, err
	}

	return coefficients, nil
}