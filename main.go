package main

import "fmt"

func main() {
	//- Khai báo biến
	// var name type
	//- Nếu không gán giá trị,
	// go sẽ tự động hán khởi tạo giá trị zero tương ứng
	// với kiểu dữ liệu của biến.
	//- Không thể thay đổi kiểu dữ liệu của biến.
	var name string // name sẽ được gán giá trị là ""
	var age int // age sẽ được gán giá trị là 0
	//- Khai báo biến có giá trị khởi tạo
	// var name type = initialvalue
	var age1 int = 29
	//- Suy luận kiểu, Go có thể tự động suy ra kiểu của biến
	// từ giá trị khởi tạo.
	var age2 = 10
	//- Khai báo nhanh
	age6 := 20
	//- Khai báo nhiều biến cùng một lúc
	var name1,name2 string
	var age3,age4 int = 30,50
	// - Khai báo các biến có các kiểu khác nhau trong cùng một câu lệnh
	var (
		name3 = "name"
		age5 = 29
	)
	name4,age7 := "naveen", 29

	fmt.Println(name,name1,name2, name3, name4, age, age1, age2, age3, age4, age5, age6, age7)
}