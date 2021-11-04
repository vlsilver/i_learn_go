package main

func main() {
	//- Kiểu bool true/false
	checkA := true
	//- Kiểu dữ liệu số (numberic type)
	//+ int8, int16, int32, int64, int
	intA := 1
	//+ float32, float64 - không chính xác
	floatA := 14.64 // Các giá trị dấy phẩy động hay số thực được go suy ra là kiểu float64
	floatB := -3.3
	floatResult := floatA / floatB
	//+ complex64, complex128 - Số phức
	complexA := 8 + 27i
	complexB := 10 + 27i
	complexResult := complexA + complexB
	//+ byte - là một cách gọi khác của uint8 [0,255]
	//+ rune - là một cách gọi khác của int32 [-2^31, 2^31-1]
	//- Kiểu string
	stringA := "vlsilver"
	//- Ngôn ngữ Go rất nghiêm ngặt và chặt chẽ, nên chúng không cho phép
	//tự đổng chuyển đổi(ép kiểu) kiểu dữ liệu.
	sum := float64(intA) + floatA
	println(checkA, intA, floatResult, complexA, complexResult, stringA, sum)


}