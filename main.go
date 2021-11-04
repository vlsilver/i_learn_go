package main

import (
	"fmt"
)

func main() {
	//Giá trị biến là hằng số và không được thay đổi.
	//Giá trị của 1 hằng số cần được biết ngay tại thời gian biên dịch.
	//Vì vậy giá trị gán cho nó cũng phải là 1 constant.
	const constA = 55
	//const constError = math.Abs(4) -- error
	//Kiểu của một hằng số là untyped(không có kiểu).
	//Hằng số chỉ cung cấp type mặc định của mình khi có dòng lệnh nào đó yêu cầu.
	//Để tạo ra một biến const có kiểu xác định ta thực hiện.
	const name = "vlsilver" // type name is untyped
	const age = 10 // type age is untyped
	const ageInt int = 10 // tạo ra hằng số ageInt với một type chỉ định
	const ageFloat float32 = age
	const ageInt32 int32 = age
	const ageComplex64 complex64 = age
	//Cú pháp của age là biểu thị chung của số, tức nó có thể là float, interger hay là 1 số phức.
	//Nên có thể gán cho nó bất kỳ kiểu dữ liệu nào.
	fmt.Printf("type of name %T", name) // Hằng số name sẽ cung cấp type mặc định string của mình
	fmt.Printf("\ntype of age %T", age) // Hằng số age cung cấp type mặc định int của mình
	type myString string
	type myInt int
	const nameTets myString = name
		// vì name là hằng số chưa chỉ định type nên đang là untyped và có thể gán vào myString
	const ageTest myInt = age
		// vì age là hằng số chưa chỉ định type nên đang là untyped và có thể gán vào myInt
	//const ageTestInt myInt = ageInt
		// vì ageInt đã chỉ định là int nên không thể gán vào myInt


}