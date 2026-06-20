package main

import (
	"fmt"
	// "math"
	// "ppt/ch5"
	"log"
	// "ppt/ch1"
	"ppt/utils"
	// "ppt/ch2"
	"ppt/ch3"
	// "ppt/ch6"
	// "ppt/ch4"
	// "ppt/ch7"
)

func main() {
	// // ---------------- CHOLESKY DECOMPOSITION -----------------------------
	// // Khai báo ma trận đối xứng xác định dương A
	// A := [][]float64{
	// 	{4, 12, -16},
	// 	{12, 37, -43},
	// 	{-16, -43, 98},
	// }

	// fmt.Println("=== Thực thi thuật toán phân rã Cholesky ===")
	// utils.InMaTran("Ma trận gốc A:", A)

	// // Chạy thuật toán
	// L, err := ch3.CholeskyDecomposition(A)
	// if err != nil {
	// 	log.Fatalf("Lỗi: %v", err)
	// }

	// // Hiển thị ma trận tam giác dưới L thu được
	// utils.InMaTran("Ma trận tam giác dưới L tìm được:", L)





	// // Thử nghiệm trực quan hóa số thực 0.15625 (Một số có thể biểu diễn chính xác ở hệ nhị phân)
	// ch1.IEEEVisualizer32(0.15625)

	// // Thử nghiệm trực quan hóa số 0.1 (Số gây ra vòng lặp vô hạn ở hệ nhị phân - dẫn đến sai số làm tròn)
	// ch1.IEEEVisualizer32(0.1)

	// // ---------------- XẤP XỈ TÍCH PHÂN BẰNG PHƯƠNG PHÁP HÌNH THANG (TRAPEZOIDAL) -----------------------------
	// a := 0.0
	// b := math.Pi
	// N := 100 // Chia thành 100 mảnh hình thang nhỏ

	// // Hàm số kiểm thử f(x) = sin(x)
	// f := func(x float64) float64 {
	// 	return math.Sin(x)
	// }

	// fmt.Println("=== Kiểm thử Phương pháp Hình thang (Trapezoidal Rule) ===")

	// // Thực thi thuật toán
	// approxResult, err := ch6.TrapezoidalIntegration(a, b, N, f)
	// if err != nil {
	// 	fmt.Println("Lỗi:", err)
	// 	return
	// }

	// exactResult := 2.0 // Kết quả giải tích chính xác

	// // Xuất kết quả đối chiếu
	// fmt.Printf("Thông số bài toán:\n")
	// fmt.Printf("  Đoạn tích phân : [%.4f, %.4f]\n", a, b)
	// fmt.Printf("  Số khoảng chia : %d\n", N)
	// fmt.Printf("  Bước lưới h    : %.6f\n", (b-a)/float64(N))

	// fmt.Printf("\nKết quả phân tích:\n")
	// fmt.Printf("  Giá trị xấp xỉ số (Hình thang) : %.6f\n", approxResult)
	// fmt.Printf("  Giá trị chính xác giải tích    : %.6f\n", exactResult)
	// fmt.Printf("  Sai số tuyệt đối thực tế       : %e\n", math.Abs(exactResult-approxResult))






	// // ---------------- XẤP XỈ TÍCH PHÂN BẰNG PHƯƠNG PHÁP HÌNH CHỮ NHẬT GIỮA (MIDPOINT) -----------------------------
	// a := 0.0
	// b := 3.0
	// N := 100 // Chia thành 100 hình chữ nhật nhỏ

	// // Hàm số kiểm thử f(x) = x^2
	// f := func(x float64) float64 {
	// 	return x * x
	// }

	// fmt.Println("=== Kiểm thử Phương pháp Hình chữ nhật giữa (Midpoint) ===")

	// // Thực thi thuật toán
	// approxResult, err := ch6.MidpointRectangleIntegration(a, b, N, f)
	// if err != nil {
	// 	fmt.Println("Lỗi:", err)
	// 	return
	// }

	// exactResult := 9.0 // Kết quả giải tích chính xác

	// // Xuất kết quả đối chiếu
	// fmt.Printf("Thông số bài toán:\n")
	// fmt.Printf("  Đoạn tích phân : [%.1f, %.1f]\n", a, b)
	// fmt.Printf("  Số khoảng chia : %d\n", N)
	// fmt.Printf("  Bước lưới h    : %.4f\n", (b-a)/float64(N))

	// fmt.Printf("\nKết quả phân tích:\n")
	// fmt.Printf("  Giá trị xấp xỉ số (Midpoint) : %.6f\n", approxResult)
	// fmt.Printf("  Giá trị chính xác giải tích  : %.6f\n", exactResult)
	// fmt.Printf("  Sai số tuyệt đối thực tế     : %e\n", math.Abs(exactResult-approxResult))







	// // ---------------- XẤP XỈ ĐẠO HÀM BẬC HAI TẠI ĐIỂM NÚT TRUNG TÂM BẰNG CÔNG THỨC ĐIỂM GIỮA -----------------------------
	// x0 := 0.0
	// h := 0.1

	// // Hàm số kiểm thử f(x) = cos(x)
	// f := func(x float64) float64 {
	// 	return math.Cos(x)
	// }

	// // Thu thập dữ liệu tại 3 điểm nút cách đều đối xứng qua x0
	// fMinusH := f(x0 - h) // cos(-0.1)
	// f0      := f(x0)     // cos(0.0)
	// fPlusH  := f(x0 + h) // cos(0.1)

	// fmt.Println("=== Kiểm thử Công thức đạo hàm bậc hai điểm giữa ===")

	// // Thực thi thuật toán xấp xỉ đạo hàm bậc hai
	// approxDeriv, err := ch6.SecondDerivMidpoint(h, fMinusH, f0, fPlusH)
	// if err != nil {
	// 	fmt.Println("Lỗi:", err)
	// 	return
	// }

	// // Đạo hàm bậc hai chính xác theo giải tích: f''(x) = -cos(x)
	// exactDeriv := -math.Cos(x0)

	// // Xuất kết quả đối chiếu
	// fmt.Printf("Giá trị các điểm nút đầu vào (h = %.1f):\n", h)
	// fmt.Printf("  f(x0 - h) = f(-0.1) = %.6f\n", fMinusH)
	// fmt.Printf("  f(x0)     = f(0.0)  = %.6f\n", f0)
	// fmt.Printf("  f(x0 + h) = f(0.1)  = %.6f\n", fPlusH)

	// fmt.Printf("\nKết quả phân tích:\n")
	// fmt.Printf("  Đạo hàm bậc hai xấp xỉ : %.6f\n", approxDeriv)
	// fmt.Printf("  Đạo hàm bậc hai chính xác: %.6f\n", exactDeriv)
	// fmt.Printf("  Sai số tuyệt đối thực tế : %e\n", math.Abs(exactDeriv-approxDeriv))


	// // ---------------- XẤP XỈ ĐẠO HÀM BẬC NHẤT TẠI ĐIỂM NÚT BIÊN BẰNG CÔNG THỨC NĂM ĐIỂM CỰC TRỊ -----------------------------
	// x0 := 2.0
	// h := 0.1

	// // Hàm số kiểm thử f(x) = x * e^x
	// f := func(x float64) float64 {
	// 	return x * math.Exp(x)
	// }

	// // Thu thập dữ liệu tại 5 điểm nút liên tiếp bắt đầu từ biên x0
	// f0 := f(x0)         // f(2.0)
	// f1 := f(x0 + h)     // f(2.1)
	// f2 := f(x0 + 2.0*h) // f(2.2)
	// f3 := f(x0 + 3.0*h) // f(2.3)
	// f4 := f(x0 + 4.0*h) // f(2.4)

	// fmt.Println("=== Kiểm thử Công thức năm điểm cực trị (Five-Point Endpoint) ===")

	// // Thực thi thuật toán xấp xỉ đạo hàm tại điểm biên
	// approxDeriv, err := ch6.FivePointEndpoint(h, f0, f1, f2, f3, f4)
	// if err != nil {
	// 	fmt.Println("Lỗi:", err)
	// 	return
	// }

	// // Đạo hàm chính xác theo giải tích: f'(2.0) = (2 + 1) * e^2
	// exactDeriv := (x0 + 1.0) * math.Exp(x0)

	// // Xuất kết quả đối chiếu
	// fmt.Printf("Giá trị các điểm nút đầu vào (h = %.1f):\n", h)
	// fmt.Printf("  f(2.0) = %.6f\n", f0)
	// fmt.Printf("  f(2.1) = %.6f\n", f1)
	// fmt.Printf("  f(2.2) = %.6f\n", f2)
	// fmt.Printf("  f(2.3) = %.6f\n", f3)
	// fmt.Printf("  f(2.4) = %.6f\n", f4)

	// fmt.Printf("\nKết quả phân tích:\n")
	// fmt.Printf("  Đạo hàm xấp xỉ (Five-Point Endpoint) : %.6f\n", approxDeriv)
	// fmt.Printf("  Đạo hàm chính xác (Giải tích)        : %.6f\n", exactDeriv)
	// fmt.Printf("  Sai số tuyệt đối thực tế             : %e\n", math.Abs(exactDeriv-approxDeriv))








	// // ---------------- XẤP XỈ ĐẠO HÀM BẬC NHẤT TẠI ĐIỂM NÚT TRUNG TÂM BẰNG CÔNG THỨC NĂM ĐIỂM TRUNG TÂM -----------------------------
	// x0 := 2.0
	// h := 0.1

	// // Hàm số kiểm thử f(x) = x * e^x
	// f := func(x float64) float64 {
	// 	return x * math.Exp(x)
	// }

	// // Thu thập giá trị tại 4 điểm nút đối xứng xung quanh x0
	// fMinus2h := f(x0 - 2.0*h) // f(1.8)
	// fMinusH  := f(x0 - h)     // f(1.9)
	// fPlusH   := f(x0 + h)     // f(2.1)
	// fPlus2h  := f(x0 + 2.0*h) // f(2.2)

	// fmt.Println("=== Kiểm thử Công thức năm điểm trung tâm (Five-Point Midpoint) ===")

	// // Thực thi thuật toán xấp xỉ đạo hàm
	// approxDeriv, err := ch6.FivePointMidpoint(h, fMinus2h, fMinusH, fPlusH, fPlus2h)
	// if err != nil {
	// 	fmt.Println("Lỗi:", err)
	// 	return
	// }

	// // Đạo hàm chính xác theo giải tích: f'(x) = (x + 1) * e^x
	// exactDeriv := (x0 + 1.0) * math.Exp(x0)

	// // Xuất kết quả đối chiếu
	// fmt.Printf("Giá trị các điểm nút đầu vào (h = %.1f):\n", h)
	// fmt.Printf("  f(x0 - 2h) = f(1.8) = %.6f\n", fMinus2h)
	// fmt.Printf("  f(x0 - h)  = f(1.9) = %.6f\n", fMinusH)
	// fmt.Printf("  f(x0 + h)  = f(2.1) = %.6f\n", fPlusH)
	// fmt.Printf("  f(x0 + 2h) = f(2.2) = %.6f\n", fPlus2h)

	// fmt.Printf("\nKết quả phân tích:\n")
	// fmt.Printf("  Đạo hàm xấp xỉ (Five-Point) : %.6f\n", approxDeriv)
	// fmt.Printf("  Đạo hàm chính xác (Giải tích): %.6f\n", exactDeriv)
	// fmt.Printf("  Sai số tuyệt đối thực tế     : %e\n", math.Abs(exactDeriv-approxDeriv))








	// // ---------------- XẤP XỈ ĐẠO HÀM BẬC NHẤT TẠI ĐIỂM NÚT TRUNG TÂM BẰNG CÔNG THỨC BA ĐIỂM TRUNG TÂM -----------------------------
	// x0 := 0.5
	// h := 0.1

	// // Hàm số kiểm thử f(x) = sin(x)
	// f := func(x float64) float64 {
	// 	return math.Sin(x)
	// }

	// // Thu thập giá trị tại các điểm nút đối xứng
	// fMinusH := f(x0 - h) // sin(0.4)
	// fPlusH := f(x0 + h)  // sin(0.6)

	// fmt.Println("=== Kiểm thử Công thức ba điểm trung tâm (Midpoint) ===")

	// // Thực thi tính đạo hàm xấp xỉ
	// approxDeriv, err := ch6.ThreePointMidpoint(h, fMinusH, fPlusH)
	// if err != nil {
	// 	fmt.Println("Lỗi:", err)
	// 	return
	// }

	// // Đạo hàm chính xác theo giải tích: d(sin(x))/dx = cos(x) -> cos(0.5)
	// exactDeriv := math.Cos(x0)

	// // Xuất kết quả
	// fmt.Printf("Giá trị các điểm nút đầu vào:\n")
	// fmt.Printf("  f(x0 - h) = f(%.1f) = %.6f\n", x0-h, fMinusH)
	// fmt.Printf("  f(x0 + h) = f(%.1f) = %.6f\n", x0+h, fPlusH)

	// fmt.Printf("\nKết quả phân tích:\n")
	// fmt.Printf("  Đạo hàm xấp xỉ (Midpoint) : %.6f\n", approxDeriv)
	// fmt.Printf("  Đạo hàm chính xác (Cos)     : %.6f\n", exactDeriv)
	// fmt.Printf("  Sai số tuyệt đối thực tế    : %e\n", math.Abs(exactDeriv-approxDeriv))




	// // ----------- XẤP XỈ ĐẠO HÀM BẬC NHẤT TẠI BIÊN BẰNG CÔNG THỨC BA ĐIỂM ĐIỂM NÚT -----------------------------
	// // Điểm nút cần tính đạo hàm ở biên
	// x0 := 0.0
	// h := 0.1

	// // Thiết lập hàm số f(x) = sin(x)
	// f := func(x float64) float64 {
	// 	return math.Sin(x)
	// }

	// // Thu thập giá trị tại 3 điểm nút cách đều
	// f0 := f(x0)        // sin(0.0)
	// f1 := f(x0 + h)    // sin(0.1)
	// f2 := f(x0 + 2*h)  // sin(0.2)

	// fmt.Println("=== Kiểm thử Công thức ba điểm cực trị ===")
	
	// // Chạy thuật toán xấp xỉ đạo hàm
	// approxDeriv, err := ch6.ThreePointEndpoint(h, f0, f1, f2)
	// if err != nil {
	// 	fmt.Println("Lỗi:", err)
	// 	return
	// }

	// // Giá trị đạo hàm lý thuyết: d(sin(x))/dx = cos(x) tại x = 0 là cos(0) = 1
	// exactDeriv := math.Cos(x0)

	// // Xuất kết quả kiểm tra
	// fmt.Printf("Giá trị các điểm nút đầu vào:\n")
	// fmt.Printf("  f(%.1f) = %.6f\n", x0, f0)
	// fmt.Printf("  f(%.1f) = %.6f\n", x0+h, f1)
	// fmt.Printf("  f(%.1f) = %.6f\n", x0+2*h, f2)
	
	// fmt.Printf("\nKết quả phân tích:\n")
	// fmt.Printf("  Đạo hàm xấp xỉ (Endpoint) : %.6f\n", approxDeriv)
	// fmt.Printf("  Đạo hàm chính xác (Cos)    : %.6f\n", exactDeriv)
	// fmt.Printf("  Sai số tuyệt đối thực tế   : %e\n", math.Abs(exactDeriv-approxDeriv))




	// // ------------------- XẤP XỈ BÌNH PHƯƠNG TỐI THIỂU TRỰC GIAO ------------------------------
	// // Hàm số cần xấp xỉ: f(x) = ln(x)
	// f := func(x float64) float64 {
	// 	return math.Log(x)
	// }

	// a := 1.0
	// b := 2.0
	// degree := 3

	// fmt.Println("=== Xấp xỉ Bình phương tối thiểu Trực giao (Legendre) ===")
	// coeffs := ch5.OrthogonalLeastSquares(a, b, degree, f)

	// // In ra các hệ số c_i đại diện cho tổ hợp tuyến tính của các đa thức trực giao
	// for i, val := range coeffs {
	// 	fmt.Printf("  Hệ số c[%d] = %.6f\n", i, val)
	// }

	// // Đánh giá sai số tại điểm x = 1.5
	// xTest := 1.5
	// fActual := f(xTest)
	// fApprox := ch5.EvaluateOrthogonalPoly(a, b, coeffs, xTest)

	// fmt.Printf("\nKiểm tra tại x = %.2f:\n", xTest)
	// fmt.Printf("  Giá trị thực tế f(x)   = %.6f\n", fActual)
	// fmt.Printf("  Giá trị xấp xỉ P(x)   = %.6f\n", fApprox)
	// fmt.Printf("  Sai số tuyệt đối       = %.6f\n", math.Abs(fActual-fApprox))

	// // XẤP XỈ BÌNH PHƯƠNG TỐI THIỂU LIÊN TỤC
	// // Hàm số cần xấp xỉ: f(x) = e^x
	// f := func(x float64) float64 {
	// 	return math.Exp(x)
	// }

	// a := 0.0
	// b := 1.0
	// degree := 2 // Tìm đa thức bậc 2

	// fmt.Println("=== Xấp xỉ Bình phương tối thiểu liên tục ===")
	// coeffs, err := ch5.ContinuousLeastSquares(a, b, degree, f)
	// if err != nil {
	// 	fmt.Println("Lỗi:", err)
	// 	return
	// }

	// // In đa thức kết quả
	// fmt.Printf("Đa thức xấp xỉ bậc %d trên đoạn [%.1f, %.1f] là:\n", degree, a, b)
	// fmt.Printf("P(x) = %.4f + %.4fx + %.4fx^2\n", coeffs[0], coeffs[1], coeffs[2])

	// // Thử đánh giá sai số tại điểm giữa x = 0.5
	// xTest := 0.5
	// fActual := f(xTest)
	// fApprox := coeffs[0] + coeffs[1]*xTest + coeffs[2]*math.Pow(xTest, 2)
	// fmt.Printf("\nKiểm tra tại x = %.2f:\n", xTest)
	// fmt.Printf("  Giá trị thực tế f(x)   = %.6f\n", fActual)
	// fmt.Printf("  Giá trị xấp xỉ P(x)   = %.6f\n", fApprox)
	// fmt.Printf("  Sai số tuyệt đối       = %.6f\n", math.Abs(fActual-fApprox))
	
	// XẤP XỈ BÌNH PHƯƠNG TỐI THIỂU RỜI RẠC
	// // Tập dữ liệu thực nghiệm mẫu: (1, 2), (2, 5), (3, 5)
	// X := []float64{1.0, 2.0, 3.0}
	// Y := []float64{2.0, 5.0, 5.0}

	// fmt.Println("=== Chạy thử thuật toán Bình phương tối thiểu rời rạc ===")
	
	// res, err := ch5.LinearLeastSquares(X, Y)
	// if err != nil {
	// 	log.Fatalf("Lỗi thực thi thuật toán: %v", err)
	// }

	// // In kết quả định dạng 2 chữ số thập phân
	// fmt.Printf("Hệ số tìm được:\n")
	// fmt.Printf("  a (hệ số góc) = %.2f\n", res.A)
	// fmt.Printf("  b (hệ số chặn) = %.2f\n", res.B)
	// fmt.Printf("=> Hàm số xấp xỉ tốt nhất: y = %.2fx + %.2f\n", res.A, res.B)


	// f := func(t float64, y float64) float64 {
	// 	return y - t*t + 1.0
	// }

	// fmt.Println(ch7.Heun(0.0, 2.0, 10, 0.5, f))
	// ch7.RungeKutta4(0.0, 2.0, 10, 0.5, f)


	// points := ch7.Midpoint(0.0, 2.0, 10, 0.5, f)

	// for idx, p := range points {
	// 	fmt.Printf("Bước %d -> t = %.1f, y = %.4f\n", idx, p.T, p.W)
	// }
	
	
	// // Định nghĩa hàm vế phải f(t, y) = y - t^2 + 1
	// f := func(t float64, y float64) float64 {
	// 	return y - t*t + 1.0
	// }

	// // Các thông số đầu vào từ sách:
	// a := 0.0      // Điểm đầu
	// b := 2.0      // Điểm cuối
	// N := 10       // Số khoảng chia (N = 10 thì h = 0.2)
	// alpha := 0.5  // Điều kiện ban đầu y(0) = 0.5

	// // Chạy thuật toán Runge-Kutta 4
	// ch7.RungeKutta4(a, b, N, alpha, f)
	
	// ch7.Euler(
	// 	0.0,
	// 	2.0,
	// 	10,
	// 	0.5,
	// 	func(t, y float64) float64 {
	// 		return y - t*t + 1
	// 	},
	// )
	
	// fmt.Println(ch4.NaturalSpline(
	// 	[]float64{1,2,3},
	// 	[]float64{2,3,5},
	// ))

	// fmt.Println(ch4.NaturalSpline(
	// 	[]float64{0,1,2,3},
	// 	[]float64{1,math.E, math.E*math.E, math.E*math.E*math.E},
	// ))
	
	
	// fmt.Println(ch4.Newton(
    //     []float64{1.0, 1.3, 1.6, 1.9, 2.2},
    //     []float64{0.7651977, 0.6200860, 0.4554022, 0.2818186, 0.1103623},
    // ))


	// f := func (x float64) float64 {
	// 	return 1/x
	// }

	// fmt.Println(ch4.LinearLagrange(5,2,4,5,1))
	// fmt.Println(ch4.GeneralLagrange(
	// 	3,
	// 	[]float64 {2, 2.75, 4},
	// 	[]float64{f(2), f(2.75), f(4)},
	// 	),
	// )

	// ch3.Gauss_seidel(
	// 	4, 
	// 	[][]float64{
	// 		{10,-1,2,0},
	// 		{-1,11,-1,3},
	// 		{2,-1,10,-1},
	// 		{0,3,-1,8},
	// 	},
	// 	[]float64{6,25,-11,15},
	// 	[]float64{0,0,0,0},
	// 	1e-3,
	// 	100,
	// )

	// ch3.Jacobi(
	// 	4, 
	// 	[][]float64{
	// 		{10,-1,2,0},
	// 		{-1,11,-1,3},
	// 		{2,-1,10,-1},
	// 		{0,3,-1,8},
	// 	},
	// 	[]float64{6,25,-11,15},
	// 	[]float64{0,0,0,0},
	// 	1e-3,
	// 	11,
	// )
	
	// augMatrix := [][]float64{
	// 	{8,2,-7},
	// 	{3, 5, 2, 8},
	// 	{6,2,8,26},
	// }

	// ch3.Gauss(3, augMatrix)
	
	// TOL := 1e-8
	// N0 := 30
	// f := func(x float64) float64 {
	// 	return math.Pow(x, 100) - math.Pow(x, 21) + 6*x*x - 15*x + 1
	// }

	
	// ch2.False_position(0, 1, f, TOL, N0)
	// ch2.Secant(0, 1, f, TOL, N0)
	// ch2.Newton(0, f, TOL, N0)
	// fmt.Printf("Bisection: %v\n", ch2.Bisection(0, 1, f, TOL, N0))

	// ch2.Fixed_point(0.5, func(x float64) float64 {
	// return math.Cos(x)
	// }, TOL, N0)

	// fixed point
	// // g(x) functions from page 61
	// g1 := func(x float64) float64 {
	// 	return x - x*x*x - 4*x*x + 10
	// }

	// g2 := func(x float64) float64 {
	// 	return math.Pow(10/x - 4*x, 0.5)
	// }

	// g3 := func(x float64) float64 {
	// 	return 0.5 * math.Pow(10 - x*x*x, 0.5)
	// }

	// g4 := func(x float64) float64 {
	// 	return math.Pow(10/(x+4), 0.5)
	// }

	// g5 := func (x float64) float64 {
	// 	return x - (x*x*x + 4*x*x - 10)/(3*x*x+8*x)
	// }

	// ch2.Fixed_point(1.5, g1, 1e-100, 30)
	// ch2.Fixed_point(1.5, g2, 1e-100, 30)
	// ch2.Fixed_point(1.5, g3, 1e-100, 30)
	// ch2.Fixed_point(1.5, g4, 1e-100, 30)
	// ch2.Fixed_point(1.5, g5, 1e-100, 30)
}