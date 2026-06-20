package utils

import (
	"fmt"
	"math"
)

type Point struct {
	T float64
	W float64
}

func EuclideanDistance(x, x0 []float64) float64 {
	if len(x) != len(x0) {
		panic("Two slices must have same lengths.")
	}

	sumOfSquares := 0.0

	for i := 0; i < len(x); i++ {
		diff := x[i] - x0[i]
		sumOfSquares += diff * diff
	}

	return math.Sqrt(sumOfSquares)
}

// TargetFunc đại diện cho hàm số f(x) cần xấp xỉ
type TargetFunc func(float64) float64

// IntegrableFunc đại diện cho hàm số f(x) cần tính tích phân
type IntegrableFunc func(float64) float64

// InMaTrat hỗ trợ hiển thị ma trận đẹp mắt
func InMaTran(label string, matrix [][]float64) {
	fmt.Println(label)
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%8.2f ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}