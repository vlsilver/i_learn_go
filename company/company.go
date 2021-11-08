package company

import "fmt"

/*
   Interface khai báo hành vi của một đối tương.
   -> Tập hợp khai báo của các method.
   -> Khi một đối tượng khác định nghĩa tất cả các phương thức của interface. thì
   -> nó được gọi là implement interface đó.
   Nó chỉ xác định đối tượng có thể làm những gì.
   Cách thực hiện như thế nào thì tuỳ thuộc vào các đối tượng implement nó
	-> Khi các đối tượng mới xuất hiện có thể dễ dàng sử dụng dựa trên tính năng implement interface.
*/
type SalaryCalculator interface {
	Calculate() int
}
type Parmanent struct {
	ID       int
	BasicPay int
	Pf       int
}

/*
   Permanent implement Caculate() của SalaryCalculator Interface
   -> Parmanent là interface implement SalaruCalculator
*/
func (receiver Parmanent) Calculate() int {
	fmt.Println("Calculator Salary of Parmanent")
	return receiver.BasicPay + receiver.Pf
}

type Contract struct {
	ID       int
	BasicPay int
}

func (contract Contract) Calculate() int {
	fmt.Println("Calculator Salary of Contract")
	return contract.BasicPay
}

func CalculatorOfEmployee(employee []SalaryCalculator) int {
	totalSalary := 0
	for _, e := range employee {
		totalSalary += e.Calculate()
	}
	return totalSalary
}
