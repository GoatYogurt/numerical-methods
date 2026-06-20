package ch1

import (
	"fmt"
	"math"
)

// IEEEVisualizer32 bóc tách và in trực quan số float32 dưới dạng các chuỗi bit nhị phân
func IEEEVisualizer32(val float32) {
	// Đọc trực tiếp 32 bit trong bộ nhớ mà không đổi giá trị toán học
	bits := math.Float32bits(val)

	// 1. Bit Dấu (Sign bit) - Bit thứ 31 (tính từ 0)
	sign := bits >> 31

	// 2. Phần mũ (Exponent) - 8 bit tiếp theo (từ bit 23 đến 30)
	// Dùng mặt nạ 0xFF (11111111 nhị phân) để lọc lấy 8 bit sau khi dịch
	exponent := (bits >> 23) & 0xFF

	// 3. Phần định trị (Mantissa / Fraction) - 23 bit cuối cùng (từ bit 0 đến 22)
	// Dùng mặt nạ 0x7FFFFF (23 bit 1 nhị phân)
	mantissa := bits & 0x7FFFFF

	fmt.Printf("\n=== TRỰC QUAN HÓA SỐ: %v (float32) ===\n", val)
	fmt.Printf("Giá trị bit thô (Hệ 16): 0x%08X\n\n", bits)

	// Hiển thị dạng sơ đồ phân rã chuỗi bit
	fmt.Println("Cấu trúc phân rã IEEE 754 (32-bit):")
	fmt.Println("-----------------------------------------------------------------")
	fmt.Printf(" [Dấu] |      [Phần Mũ - 8 bit]       |    [Phần Định Trị - 23 bit]   \n")
	fmt.Println("-----------------------------------------------------------------")
	fmt.Printf("   %d   |  %08b  |  %023b\n", sign, exponent, mantissa)
	fmt.Println("-----------------------------------------------------------------")

	// Giải thích ý nghĩa toán học số trị
	fmt.Printf("\nPhân tích thành phần toán học:\n")
	fmt.Printf(" 1. Biển dấu: %d -> %s\n", sign, map[uint32]string{0: "Số Dương (+)", 1: "Số Âm (-)"}[sign])
	
	// Số thực lưu phần mũ có cộng độ lệch (bias) là 127 đối với float32
	actualExponent := int(exponent) - 127
	fmt.Printf(" 2. Phần mũ gốc lưu trong máy: %d\n", exponent)
	fmt.Printf("    => Số mũ thực tế (đã trừ độ lệch 127): 2^(%d)\n", actualExponent)
	
	// Tính giá trị phần định trị dạng thập phân
	mantissaValue := 1.0
	for i := 1; i <= 23; i++ {
		// Kiểm tra xem bit thứ (23-i) có bật hay không
		if (mantissa >> (23 - i) & 1) == 1 {
			mantissaValue += math.Pow(2, float64(-i))
		}
	}
	fmt.Printf(" 3. Phần định trị hệ 10 phục hồi: %f (Đã tự động cộng bit 1 ẩn đầu tiên)\n", mantissaValue)

	// Tái dựng lại công thức toán học gốc
	finalCalc := math.Pow(-1, float64(sign)) * mantissaValue * math.Pow(2, float64(actualExponent))
	fmt.Printf("\n=> Công thức tái tạo: (-1)^%d * %f * 2^(%d) = %g\n", sign, mantissaValue, actualExponent, finalCalc)
}