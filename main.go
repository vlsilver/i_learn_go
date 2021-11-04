//Mỗi file go đều bắt đầu với câu lệnh package name
//Package được dùng để phân chia code cho các mục đích sử dịnh khác nhau
// và tăng tính tái sử dụng, ở đây package name là name.
package main

//package fmt được import vào file.
//những funtion của fmt library có thể được sử dụng trong file

import "fmt"

//Chương trình được thực thi bắt đầu từ một hàm main() duy nhất.
func main() {
	fmt.Println("Hello Go!")
}