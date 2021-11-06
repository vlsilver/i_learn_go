package main

import "fmt"

func main() {
	employee := Employee{
		name:     "vlsilver",
		salary:   5000,
		currency: "$",
	}
	employee.displaySalary()
	employee.growSalary(10000)
	fmt.Println("Emplyee when grow salary", employee)
	employee.updateAddress("New home in here")
	fmt.Println("Employee when update Address:", employee)
	ePointer := &employee
	ePointer.showAddress()
	showAddress(employee)
	/*
		Lỗi biên dịch không thể gán con trỏ cho struct
	*/
	//showAddress(ePointer)
	fmt.Println("Employee when update Address:", employee)
	myInt(8).add(5)
}

type Employee struct {
	name     string
	salary   int
	currency string
	Address
}

type Address struct {
	address string
}

/*
 displaySalary() method
-> là một funcion được khai báo riêng cho một kiểu dữ liệu.
-> kiểu dữ liệu này được gọi là vật nhận(reciver).
-> các method có thể trùng tên nếu vật nhận khác nhau.
-> vật nhận có thể là struct hoặc non-struct và phải có sẵn
func (t Type) nameMethod() {}
*/

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

/*
growSalary() method
Reciever là con trỏ so với vật nhận là giá trị
-> Với những vật nhận là con trỏ thì sự thay đổi giá trị của reciver bên trong method sẽ dẫn đến sự thay đổi của reciever
-> còn vật nhận là giá trị thì không.
*/
func (ePointer *Employee) growSalary(salary int) {
	oldSalary := ePointer.salary
	ePointer.salary = salary
	fmt.Printf("Salary grow from $%d -> $%d\n", oldSalary, salary)
}

/*
updateAddress() method
->Có thể sử dụng trực tiếp các field/method của một struct là field ẩn danh nằm trong một struct
*/
func (ePointer *Address) updateAddress(address string) {
	ePointer.address = address
	fmt.Println("Address of Employ updated to :", address)
}

/*
showAddress() method
Reciever là một giá trị trong một phương thức so với đối số là một giá trị trong một hàm
-> Khí một func có đối số là giá trị thì nó chỉ chấp nhận đối số là giá trị.
Reciever là con trỏ trong một method so với đối số là con trỏ của func
-> Khi một method có vật nhận là giá trị nó sẽ chấp nhận cả vật nhận là con trỏ lẫn vật nhận là giá trị
-> Một func có đối số là con trỏ thì nó chỉ chấp nhận con trỏ.

-> Một method có reciever là con trỏ thí nó chấp nhận cả hai, cả giá trị và con trỏ.
*/
func (e Employee) showAddress() {
	fmt.Println("Address of Employee: ", e.address)
	e.address = "Update Address"
}
func showAddress(e Employee) {
	fmt.Println("Address of Employee: ", e.address)
}

/*
Method đối với kiểu dữ liệu non-struct
-> Để định nghĩa method của một type thì method và kiểu dữ liệu của reciever phải được
định nghĩa trên cùng một package.
-> Để sử dụng các type đã được khai báo sẵn, chúng ta cần tạo một type bí danh dựa trên type đã có đó
*/
type myInt int
func (receiver myInt) add(b myInt) myInt {
	return receiver + b
}
