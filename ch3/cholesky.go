package ch3

import (
	"errors"
	"math"
)

// CholeskyDecomposition thực hiện phân rã ma trận A thành L * L^T.
// A: Ma trận vuông đối xứng và xác định dương kích thước n x n.
// Trả về ma trận tam giác dưới L.
func CholeskyDecomposition(A [][]float64) ([][]float64, error) {
	n := len(A)
	
	// Kiểm tra điều kiện ma trận vuông ban đầu
	if n == 0 || len(A[0]) != n {
		return nil, errors.New("ma trận đầu vào phải là ma trận vuông")
	}

	// Khởi tạo ma trận kết quả L kích thước n x n chứa toàn số 0
	L := make([][]float64, n)
	for i := range L {
		L[i] = make([]float64, n)
	}

	// Tiến trình tính toán các phần tử của L
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			sum := 0.0
			for k := 0; k < j; k++ {
				sum += L[i][k] * L[j][k]
			}

			if i == j { // Tính toán phần tử trên đường chéo chính
				diff := A[i][i] - sum
				// Nếu giá trị dưới căn <= 0, ma trận không xác định dương
				if diff <= 1e-15 {
					return nil, errors.New("ma trận không xác định dương, không thể thực hiện phân rã Cholesky")
				}
				L[i][j] = math.Sqrt(diff)
			} else { // Tính toán phần tử nằm dưới đường chéo chính
				L[i][j] = (A[i][j] - sum) / L[j][j]
			}
		}
	}

	return L, nil
}