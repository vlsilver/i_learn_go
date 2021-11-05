package main

import "fmt"

func main() {
	// map là kiểu dữ liệu tham chiếu giống như slice, tuy nhiên nó không có sức chứa
	// Khởi tạo map
		//->Không truyền giá trị -> var learnMap = map[string]int
			//->Giá trị mặc định của map là nil
			//->Khi thêm item cho map thì sẽ bị lỗi runtime panic
			//->Chúng ta phải dùng hàm make để khởi tạo, khi đó bộ nhớ sẽ cấp phát vùng nhớ
	// Đọc/thêm giá trị map -> map[key]
	// Check key trong map -> value,ok := map[key]
		//-> Nếu key có tồn tại sẽ có giá trị trả về cho value và ok = true
		//-> Nếu không thì giá trị mặc định của value sẽ được trả về và ok = false
	// Chúng ta có thể dùng for cho map tương tư như slice
		//-> for key,value := range map {}
	var learnMap map[string]int
	//learnMap["vl"] = 123 - error
	learnMap = make(map[string]int)
	learnMap["vl"] = 123
	learnMapAssign := learnMap
	learnMapAssign["silver"] = 456
	fmt.Println("Learn Map with assign: ",learnMap)
	fmt.Println("Learn Map with assign: ",
		learnMapAssign)
	checkKey := "silver1"
	value,ok := learnMap[checkKey]
	if ok {
		fmt.Println("Learn Map with check key in map:")
		fmt.Printf("have %v in %v and value is %v\n",checkKey, learnMap,value)
	}else {
		fmt.Printf("no %v in %v and value default is %v\n",checkKey, learnMap,value)
	}
	for key,value := range learnMap {
		fmt.Printf("With key %v value is %v\n", key, value)
	}


}
