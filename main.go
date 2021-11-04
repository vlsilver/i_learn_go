package main

import "fmt"

func main() {
	learnArray()
	learnSlices()
}

func learnArray() {
	//Mảng trong  là tập hợp của các phần tử có cùng type.
	//Kích thước của mảng là một phần của types.
	//Các mảng có kích thước khác nhau thì là những loại khác nhau.
		//->Vì vậy mảng không thể thay đổi kích cỡ
	//Mảng trong Go là loại giá trị chứ không phải là loại tham chiếu
		//->khi chúng ta gán mảng a cho một biến mới thì một bản copy
		//->của bạn gốc sẽ được gán cho biến mới.
		//->nếu thay đổi được thực hiện trên biến mới
		//->sẽ không ảnh hưởng gì trên biến cũ
		//->khi mảng được truyền cho các hàm như các tham số
		//->chúng sẽ truyền theo giá trị và mảng ban đầu sẽ không bị thay đổi

	//Khai báo mảng - không gán giá trị
	var arr1 [3]int
	fmt.Println(arr1) //[0,0,0]
	// gán giá trị cho mảng
	arr1[0] = 12
	arr1[1] = 8
	arr1[2] = 1
	//Khai báo ngắn - shot hand declaration
	arr2 := [3]int{12,8,1}
	//Không bắt buộc tất cả các phần từ phải được gán giá trị
	arr3 := [3]int{12}
	//Khai báo bỏ qua chiều dài mảng
	arr4 := [...]int{12,8,1}
	//không thể thay đổi kích thước mảng
	arr5 :=  [...]int{1}
	//arr5 = arr4 - error
	//giá trị mảng ban đầu không bị thay đổi vì mảng là kiểu giá trị
	arr6 := arr5
	arr6[0] = 10
	//truyền arr như một param, giá trị arr được truyền đi không thay đổi
	changeArrParam(arr6)
	//Vòng lặp for với mảng
	forLoopWithArray(arr4)
	// Mảng hai chiều
	arr7 := [3][2]int{
		{0,1},
		{2,3},
		{4,5},
	}
	fmt.Println(arr1) // [12,8,1]
	fmt.Println(arr2) //[12,8,1]
	fmt.Println(arr3) //[12,0,0]
	fmt.Println(arr4) //[12,8,1]
	fmt.Println(arr5) //[1]
	fmt.Println(arr6) //[10]
	fmt.Println(arr7) //
}

func changeArrParam(arr [1]int) {
	arr[0] = 55
	fmt.Println("Mảng trong hàm thay đổi:",arr)
}

func forLoopWithArray(arr [3]int) {
	for i := 0 ; i < len(arr); i++ {
		fmt.Printf("%d th element of a is %v\n",i,arr[i])
	}
	// Nếu muốn giữ lại giá trị và bỏ qua chỉ mục
		//-> for _,v := range a {}
	for i,v := range arr {
		fmt.Printf("%v th element of a is %v\n",i,v)
	}
 }

 func learnSlices() {
	 //Slice có phần tử kiểu T và được biểu thị bằng []T
	 //Slice là kiểu tham chiếu giá trị
	 //Tạo slice (từ mảng, từ slice khác -> phụ thuộc, tạo nhanh, tạo bằng hàm make, tạo bằng hàm copy)
	 //Tạo slice từ array
	 	//->slice chỉ chứa các tham chiếu đến các phần tử của mảng/slice khác
	 	//->khi phần tử trên slice thay đổi thì chính phần tử đó trong array sẽ bị thay đổi
	 //Kích thước và sức chứa của một slice có thể thay đổi
	 	//->Sức chứa của 1 slice được khởi tạo từ 1 arr/slice chính là tổng số lượng phần từ của
	 	//mảng trừ đi chỉ mục đầu tiền của slice. Bất kỳ thay đổi nào trong khoảng thay đổi này đều dẫn đến sự thay
	 	//đổi của mảng/slice gốc.
	 //Thêm phần tử cho slice bằng hàm append
	 	//->Sức chứa của slice sẽ tăng gấp đôi khi số lượng phần từ vượt quá sức chứa trước đó.
	 	//->Khi đó một slice mới sẽ được tạo ra thay thế cho slice cũ và các phần tử mới không còn làn tham chiều
	 	//đến mảng/slice gốc nữa mà là một bảng copy. Bất kỳ sự thay đổi nào cũng ko làm ảnh hưởng đến mảng/slice gốc.
	 //Thêm một slice này đến 1 slice khác bằng toán tử ...
	 //Chúng ta có thể tạo ra một slice mới dựa trên hàm copy mà không cần còn tham chiếu đến mảng/slice cơ sở

	 arr1 := [5]int{1,2,3,4,5}
	 slice1 := arr1[1:4] // len(slice1) = 3, cap(slice) = len(arr1) - 1 = 4
	 slice1[0] = 10 //arr1[0] cũng bị thay đổi theo
	 slice1 = append(slice1, 50) //slice[3] = 50 -> arr1[4] = 50
	 slice1 = append(slice1, 10) //len(slice1) > cap(slice1Old) -> cap(slice) = 2 * cap(sliceOld)
	 slice1[1] = 10 // arr[2] = 3
	 // Tạo nhanh slice
	 slice2 := []int{1,2,3,4,5}
	 //Tạo slice bằng make([]T,len,cap)
	 slice3 := make([]int,5,10)
	 slice4 := append(slice2[1:4], slice3[0:1]...)
	 //->slice 4 sẽ tham chiếu đến phần tử của slice2 vì số lượng phần tử vẫn chưa vượt quá sức chứa của nó.
	 //-> bất kì thay đổi nào của slice 4 sẽ dẫn đến thay đổi của slice2
	 //-> cap(slice4) = 4
	 //-> len(slice4) = 4
	 slice4[1] = 100 // -> slice[2] = 100
	 slice4[3] = 111 // -> slice[4] = 111
	 // tạo một bản slice copy
	 slice4Copy := make([]int, len(slice4))
	 copy(slice4Copy,slice4 )
	 slice4Copy[0] = 200 //-> slice4[0] không thay đổi

	 updateSlice(slice4)
	 fmt.Printf("type %T\n",slice1)
	 fmt.Println(arr1) //[1,10,3,4,50]
	 fmt.Println("slice1",slice1)//[10,10,4,50,10]
	 fmt.Println("slice2", slice2)//[1,2,100,4,111]
	 fmt.Println("slice3", slice3)//[0,0,0,0,0]
	 fmt.Println("slice4", slice4)//[2,100,4,111]
	 fmt.Println("slice4Copy", slice4Copy)//[2,100,4,111]
 }

 func updateSlice(slice []int) {
	 // slice được truyền vào cộng thêm phần từ và có len(slice) > cap(slice) nên không còn liên quan gì đến slice bên ngoài.
	 slice = append(slice, 55)
	 slice[3] = 101 // slice bên ngoài không thay đổi
	 fmt.Println("slice5 in update function", slice)
 }