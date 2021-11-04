package main

import "fmt"

func main() {

	name := "vlsilver"
	age := 17
	fmt.Print("hello, ")
	// Create newline with \n
	fmt.Print("vlsilver!\n")
	fmt.Print("new line \n")

	// hàm fmt.Println() đã tích hợp sẵn \n
	fmt.Println("newline")
	fmt.Println("my age is",age, "and my name is",name)

	// hàm fmt.Printf(formatted strings)
	// %_ = format specifier
	// %v truyền agrument
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	// %q truyền agrument string với "agrument"
		// nếu không phải string sẽ hiển thị sai
	fmt.Printf("my age is %q and my name is %q\n",age, name)
	// %T lấy type của biến
	fmt.Printf("age is of type %T\n", age)
	// %f show float, %0.1f round float with 1 decimal
	fmt.Printf("your scored %f point!\n", 225.55)
	fmt.Printf("your scored %0.1f point!\n", 225.55)
	// save formatted string
	str := fmt.Sprintf("my age is %v and my name is %q\n",age, name)
	fmt.Println("save string formatted:",str)

}