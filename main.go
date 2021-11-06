package main

import "fmt"

func main() {
	//Được định nghĩa cho một tập hợp các trường
	//Được sử dụng để nhóm dữ liệu vào một cấu trúc
		//-> thay vì thiết kế dưới dạng riêng biệt
	newStruct := learStruct("pie")
	fmt.Println("Create News Struct: ", newStruct)
	fmt.Println("Learn about Promoted Field:",newStruct.key)
}