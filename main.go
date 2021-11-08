package main

import (
	"fmt"
	"time"
)

/*
Hàm main sẽ chạy trong Goroutines của riêng nó gọi là main Goroutines
*/
func main() {
	go goroutines1()
	fmt.Println("Main function!")
	go numbers()
	go alphabets()
	time.Sleep(3000*time.Millisecond)
	fmt.Println("Main terminated")
}

/*
Go là ngôn ngữ lập trình đồng thời, không phải ngôn ngữ lập trình song song.
-> Concurrency là khả năng thực hiện nhiều tác vụ với nhau một cách xen kẽ.
*/

/*
Khởi tạo một Goroutines
-> Nhập từ khoá go vào trước function hoặc method.
*/
func goroutines1() {
	fmt.Println("This is a Goroutines!")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}

func alphabets() {
	alphabets := []string{"a", "b", "c", "d", "e"}
	for _, value := range alphabets{
		time.Sleep(400* time.Millisecond)
		fmt.Printf("%v\n",value)
	}
}
