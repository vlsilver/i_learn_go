package main

import "fmt"

func main() {
	//Khai báo con trỏ
		//->*T là kiểu biến con trỏ mà trỏ đến một giá trị kiểu T.
		//->var a *string = &"vlsilver"
		//->toán tử & được sử dụng để lấy địa chỉ của một biến.
		//->giá trị default của một con trỏ là nil
	//Tham chiếu ngược
		//->Truy cập vào giá trị của biến mà con trỏ đang trỏ tới.
		//->var value = *cursor
	//Khi dùng con trỏ để thay đổi giá trị bộ nhớ con trỏ đang chỉ tới
		//-> giá trị của biến cũng thay đổi
	//Go không hỗ trợ các phép toán trên con trỏ
	learnPointer1()
	learnPointer2()
	learnPointer3()
	learnPointer4()
}

// learnPointer1 khai báo
func learnPointer1()  {
	fmt.Println("learnPointer1 khai báo")
	b := 255
	var a *int = &b //a được trỏ đến b
	fmt.Println("Địa chỉ ô nhớ của b: ",a)
	fmt.Println("Địa chỉ ô nhớ của a: ",&a)
	var a1 *int
	fmt.Println("Giá trị default của con trỏ: ", a1)
	a1 = &b
	fmt.Println("Địa chỉ ô nhớ của b: ",a)
}

// learnPointer2 tham chiếu ngược
func learnPointer2()  {
	fmt.Println("learnPointer2 tham chiếu ngược")
	b := 100
	bPointer := &b
	fmt.Println("Địa chỉ ô nhớ của b: ",bPointer)
	valueB := *bPointer
	fmt.Println("Giá trị tham chiếu ngược từ pointer: ",valueB)
}

// learnPointer3 thay đổi giá trị của ô nhớ dựa trên con trỏ
func learnPointer3()  {
	fmt.Println("learnPointer3 thay đổi giá trị của ô nhớ dựa trên con trỏ")
	b := 1
	*(&b)++
	fmt.Println("Giá trị của ô nhớ sau khi dùng con trỏ thay đổi: ", b)
}

// learnPointer4 sử dụng con trỏ trong hàm
func learnPointer4()  {
	fmt.Println("learnPointer4 sử dụng con trỏ trong hàm")
	a := 58
	arr := [3]int{1,2,3}
	fmt.Println("Value của a trước khi gọi hàm change: ", a)
	fmt.Println("Value của array trước khi gọi hàm change: ", arr)
	change(&a)
	changeArr(&arr)
	fmt.Println("Value của a sau khi gọi hàm change: ", a)
	fmt.Println("Value của arr sau khi gọi hàm change: ", arr)
}
func change(val *int)  {
	*val = 55
}
func changeArr(pointer *[3]int) {
	(*pointer)[0] = 2
}