package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	learnLoop1()
	learnLoop2()
	learnLoop3()
	learnLoop4()
	learnLoop5()
	learnLoop6()
	learnLoop7()
}

func learnLoop1() {
	for i := 0; i <= 10; i++ {
		println("Learn loop1 of i: ", i)
	}
}

// learnLoop2 break keywork nếu gặp keywork này thì chương trình sẽ
//-> sẽ thoát vòng lặp for đang chạy
func learnLoop2() {
	for i := 0; i <= 10; i++ {
		if i > 5 {
			println("Learn loop2 will break at: ", i)
			break
		}
		println("Learn loop2 of i: ", i)
	}
}

// learnLoop3 continue keywork gặp keywork này chương trình sẽ bỏ qua các dòng lệnh còn lại
func learnLoop3() {
	for i := 0; i <= 10; i++ {
		if i == 3 {
			println("Learn loop3 will continue at: ", i)
			continue
		}
		a := i
		println("Learn loop3 will have a with value: ", a)
	}
}

// learnLoop4 rút gọn for
func learnLoop4() {
	i := 0
	for i <= 10 {
		println("Learn loop4 with i:", i)
		i++
	}
}

// learnLoop5 khai báo nhiều biến
func learnLoop5() {
	for i, j := 0, 3; i != 10; i = j + 1 {
		j++
		fmt.Printf("Learn loop5 will have a with value: i-%v j-%v\n", i, j)
	}
}

// learnLoop6 vòng lặp vô hạn
func learnLoop6() {
	var i int
	for {
		if i > 10 {
			break
		}
		i++
		fmt.Println("Learn loop6 will have a with value:", i)
	}
}

// learnLoop7 vòng lặp với array, slice
func learnLoop7() {
	slice := []int{412, 3412, 4343, 554, 434, 54}
	//standard
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Learn Loop7 with i:%v,slice element %v\n", i, slice[i])
	}
	//rút gọn
	for i, value := range slice {
		fmt.Printf("Learn Loop7 with i:%v,slice element %v\n", i, value)
	}
	//không dùng một trong 2
	for _, value := range slice {
		fmt.Printf("Learn Loop7 with slice element %v\n", value)
	}
}
