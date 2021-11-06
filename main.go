package main

import (
	"fmt"
	"i_go/company"
)

func main() {
	var calculator company.SalaryCalculator
	parmanent := company.Parmanent{
		BasicPay: 1000,
		Pf:       100,
	}
	contract := company.Contract{
		BasicPay: 500,
	}
	/*
		Vì Parmanent là implement SalaryCalculator nên ta có thể gán giá trị
	*/
	calculator = parmanent
	totalSalary := company.CalculatorOfEmployee([]company.SalaryCalculator{parmanent, contract})
	fmt.Println("Interface of SalaryCalculator", calculator)
	fmt.Println("Total Salary of Employee: ", totalSalary)
	describe(55)
	describe(parmanent)
	assert(56)
	safeAssert("vlsilver")
	findType("vlsilver")
	findEmployee(parmanent)
}

/*
Interface rỗng
-> Một interface không có method thì được gọi là interface rỗng.
-> Nó được biểu diễn dưới dạng interface{}
-> Mặc định tất cả các type dữ liệu đều là implement interface rỗng.
*/
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

/*
Type Asserttion - Eps kiểu giá trị cơ bản của interface
-> Cú pháp: i.(T) ép giá trị cơ bản của interface i về một kiểu cụ thể là T.
-> Nếu kiểu ép không đúng thì chương trình sẽ bị lỗi runtime panic
	-> Sử dụng cú pháp: v,ok := i.(T) để ép kiểu
*/
func assert(i interface{}) {
	s := i.(int)
	fmt.Printf("Type = %T, value = %v\n", s, s)
}
func safeAssert(i interface{}) {
	v, ok := i.(int)
	if ok {
		fmt.Printf("Type = %T, value = %v\n", v, v)
	} else {
		fmt.Printf("interface is %T not is int, value default %v\n", i, v)
	}
}
/*
Sử dụng Switch với type
 */
func findType(i interface{})  {
	switch i.(type) {
	case string:
		fmt.Printf("This is %T and value is %v\n",i.(string),i)
	case int:
		fmt.Printf("This is %T and value is %v\n",i.(int),i)
	default:
		fmt.Println("Unknow type\n")
	}
}
func findEmployee(i interface{})  {
	switch v := i.(type) {
	case company.Parmanent:
		fmt.Printf("This is %T and basic Salary is: %v\n",v,v.BasicPay)
	case company.Contract:
		fmt.Printf("This is %T and basic Salary is: %v\n",v,v.BasicPay)
	default:
		fmt.Println("Unknow Employee\n")
	}

}