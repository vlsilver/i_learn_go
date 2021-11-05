package main

import (
	"fmt"
	"math"
)

func main() {
	function1("vlsilver")
	function2([]string{"vlsilver", "tifa", "barret"}, function1)
	fmt.Printf("Diện tích của hình tròn %0.2f\n", function3(10))
	result := learnVariadicFunction(10,[]int{1,2,3,4,5,10}...)
	fmt.Println("Learn Variadic Function with result: ",result)
	learnSlice()


}

// function1 standard function
func function1(n string) {
	fmt.Printf("Learn function 1 with %v\n", n)
}

// function2 with paramater is function
func function2(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

// function with return function
func function3(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

//function với nhiều param cùng type
func function4(name, class string, age, number int) {
}
 //function với nhiều giá trị trả về
 // a,b,c := function5("vlsilver")
 func function5(name string)(int,int,int) {
	 return 1,2,3
 }
 //function với giá trị return được đặt tên
 func fucntion6(name string)(nameA,nameB string) {
	 nameA = "vl"
	 nameB = "silver"
	 return
 }
 //hàm bất định
 	 //->func variadicFunction(elements ...Type)
	 //->chấp nhận một dãy các agruments có kiểu Type
	 //->paramater cuối cùng của hàm mới được coi là bất định
 func learnVariadicFunction(num int, nums ...int)bool {
	 fmt.Println("Learn Variadic Function")
	 for _,v := range nums {
		 if v == num {
			 return true
		 }
	 }
	 return false
 }

 func learnVariadicFunctionWithSlice(s ...string) {
	 s[0] = "vl"
	 s = append(s, "nguyen")
	 s[1] = "silver"
	 fmt.Println("Learn Variadic Function with Slice: ",s)
 }

 func learnSlice() {
	 s := []string{"a","b"}
	 learnVariadicFunctionWithSlice(s...)
	 fmt.Println("Learn Slice with: ", s)
 }