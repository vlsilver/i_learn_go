package main

import "fmt"

//Khai báo struct
//Nếu là tên viết hoa chữ cái đầu tiên thì khi khai báo một struct
	//-> tạo ra một kiểu mới.
	//-> nếu là chữ thường thì không có kiểu mới được tạo ra
	//-> và được gọi là cấu trúc ẩn danh.
	//-> nếu không có tên trường mà chỉ có type thì trường đó gọi là trường ẩn danh và
	//-> khi đó tên trường được chính là type.
//Khi khai báo tạo một struct mới chúng ta có thể thay đổi ví trí các trường
	//-> đối với cách tạo mới bằng tên trường, hoặc có thể không dùng tên trường nhưng đảm bảo đúng vị trí.
//Có thể không cần xác định các giá trị của trường trong lúc tạo struct mới
	//-> giá trị của các trường sẽ là default value
	//Go cho phép ta truy cập các trường thông qua con trỏ một cách đơn giản.
	// -> structPointer := &newstruct
	// -> a := structPointer.field1 hoặc a := (*structPointer).field1
	//Chúng ta có thể tạo ra các cấu trúc lồng nhau.
	//Các trường được khuyến khích
		//->Các trường của một struct được khai báo ẩn danh trong một struct khác.
		//->các trường đó được gọi là các trường khuyến khích và được sử dựng như các trường của struct cha.
		//Nếu muốn các pakage khác truy cập được thì chúng ta cần phải viết hoa cả chữ cái đầu của Struct và
		//-> chữ hoa của các trường muốn được truy cập.


type bill struct {
	name  string
	items map[string]float64
	tip   float64
	string
	Collect
}

type Collect struct {
	key string  //Promoted fields
	value float64 //Promoted fields
}

func learStruct(name string)bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
		string:	"something here",
	}
	fmt.Println("Learn Struct make new struct: ", b)
	return b
}
