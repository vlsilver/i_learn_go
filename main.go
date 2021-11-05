

package main

import (
	"fmt"
	"i_go/rectangle"
	_"i_go/test_identifier" // import pacakage nhưng không sử dụng bất kì hàm nào
	"log"
)
// init hàm sẽ được gọi sau khi biến cục bộ được khởi tạo
// check nếu data cục bộ không thoả yêu cầu sẽ dừng chương trình.
func init() {
	fmt.Println("Hàm init() của package main được gọi")
	if len < 0 || wid < 0 {
		log.Fatalln("len and wid is less than zero")
	}
}
// biến cục bộ của package được khởi tạo đầu tiên khi package được import
var len = 10
var wid = -1

func main() {
	//packages sinh ra để chúng ta có thể tái sử dụng code hay bảo trì code.
	//chúng ta chia code ra theo các folder/moudle riêng lẻ.
	//chúng ta định danh chúng bằng packages
	//khi sử dụng chúng ta import nó.

	//Để chỉ định rằng một file mã nguồn cụ thể thuộc về gói nào đó
	//ta dùng câu lệnh package packagename.
		//-> và đây là dòng lệnh đầu tiên của mọi file mã nguồn.

	//Với các package tuỳ chỉnh ta import bằng cú pháp -> import path
	//Chỉ những hàm/biến trong packages được viết hoa chữ cái đầu tiên thì mới có thể truy cập
	//Những method và properties này được gọi là exported

	//Mỗi package có thể chứa một hàm tên là init
		//->Không có paramater, return value -> func init(){}
		//->Không được gọi một cách tường minh.
		//->Dùng để thực hiện các nhiệm vụ khởi tạo, xác minh tính chí nh xác của chương tình khi bắt đầu thực hiện
		//->Một packages có thể có nhiều hàm init - được define trong 1 fine hoặc trong nhiều file
			//->chúng được gọi theo thứ tự biên dịch
		//->Thứ tự khởi tạo của một package: biến toàn cục -> init()
		//->Nếu một package được import trong một package khác
			//->các packages sẽ được khởi tạo trước tiên
			//Một package sẽ được khởi tạo một lần duy nhất
		//Sử dụng blank identifier (định danh trống)
			//->Khi chúng ta import một package mà không sử dụng sẽ bị xem là không hợp lê.
			//->Nếu muốn import trước một số package mà không dùng ngay bây giờ, có thể thực hiện bằng bách.
				//->giả định sử dụng packages: -> var A = package.Functionname -> khuyến khích sử dụng
				//->chỉ dùng để đảm bảo việc khởi tạo, không cần sử dụng bất kỳ hàm, biến nào của package -> import (_ "namepackage")
	var len,width float64 = 6,7
	var diagonal = retangle.Diagonal(len,width)
	var area = retangle.Area(len,width)
	fmt.Printf("Diagonal of the rectangle: %0.2f\n", diagonal)
	fmt.Printf("Area of the rectangle: %0.2f\n", area)

}