package main

import (
	"fmt"
)

/*
Hàm main sẽ chạy trong Goroutines của riêng nó gọi là main Goroutines
*/
func main() {
	channelTestReadWrite()
	channelTestReadWriteMultiChannel()
	channelTestLock()
	channelTestWithLoop()
}

/*
Channels là các đường ống mà các Goroutines sử dụng để giao tiếp với nhau.
Mỗi channel có một loại dữ liệu được liên kết.
Giá trị default của channel là nil.
-> make(chan T) là cách khởi tạo channel
*/
func initChannel() chan int {
	fmt.Println("Create a channel")
	return make(chan int)
}

/*
Đọc dữ liệu từ channel
-> data := <- channel
*/
func readDataChannel(channel chan int) int {
	fmt.Println("Read data from channel")
	return <-channel
}

/*
Gửi dữ liệu vào channel
-> channel <- data
*/
func sendDataToChannel(channel chan int, data int) {
	fmt.Println("Send data to channel")
	channel <- data
}

/*
Mặc định thì gửi và đọc một channel sẽ bị chặn.
-> Khi dữ liệu gửi đến channel, điều khiện sẽ bị chặn trong câu lệnh gửi cho đến khi
-> một Goroutine khác đọc dữ liệu từ nó.
-> Khi đọc dữ liệu từ channel, việc đọc cũng bị chặn cho đến khi một Goroutine ghi dữ liệu vào channel đó.
	-> Chúng ta phải setup làm sao để hai việc đọc và gửi này xảy ra cùng nhau.
	-> Nếu không chương trình sẽ bị chặn và không kết thúc được.
*/

func channelTestReadWrite() {
	channelNumber := initChannel()
	go sendDataToChannel(channelNumber, 10)
	readDataChannel(channelNumber)
	fmt.Println("Main function End")
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func channelTestReadWriteMultiChannel() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output: ", squares+cubes)
}

/*
-> Khoá chết
	-> Nếu một Goroutine đang gửi dữ liệu trên một channel thì dự kiến một sô Goroutine khác sẽ nhận được dữ liệu.
	-> Nếu điều này không xảy ra, thì chương trình sẽ bị lỗi Deadlock.
	-> Điều tương tự xảy ra với lúc với một channel đang chờ nhận dữ liệu.
	Không có Goroutine náo read dữ liệu nên chương trình khi chạy sẽ báo lỗi	 deadlock
*/
func channelLockKey() {
	channel := make(chan int)
	channel <- 5
}

/*
Channel một chiều
-> Đó là channel chỉ gửi hoặc nhận dữ liệu.
-> Nếu chúng ta cố gắng đọc hoặc gửi dữ liệu vào một channel chỉ gửi hoặc chỉ nhận thì sẽ bị lỗi.
*/

func sendDataOneLine(sendch chan<- int) {
	sendch <- 10
}

func channelTestOneLine() {
	/*
		Tạo ra một channel chỉ send dữ liệu
	*/
	sendch := make(chan<- int)
	go sendDataOneLine(sendch)
	// fmt.Println(<-sendch) // Error
}

/*
Đóng channel và chạy các vòng lặp trên các channel
-> Người gửi có khả năng đóng channel để thông báo cho người nhận rằng sẽ không có thêm dữ liệu nào được gửi trên channel.
	-> v, ok := <- ch
	-> oke = true nếu send dữ liệu vào channel thành công.
	-> oke = false có nghĩa là chúng ta đang đọc dữ liệu từ một channel kín.
		-> Channel đó đã bị đóng.
		-> Giá trị của một channel kín là giá trị default type của channel.
*/
func producer(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

/*
Vòng lặp đọc dữ liệu từ channel đến khi nó được đóng lại, vòng lặp sẽ tự động thoát.
*/
func channelTestLock() {
	channelLock := make(chan int)
	go producer(channelLock)
	for v := range channelLock {
		fmt.Println("Recieved: ", v)
	}
}

func digits(number int, channelDigitSquare chan int, channelDigitCubes chan int) {
	for number != 0 {
		digit := number % 10
		channelDigitSquare <- digit
		channelDigitCubes <- digit
		number /= 10
	}
	close(channelDigitSquare)
	close(channelDigitCubes)
}

func calcSquaresWithLoopChannel(channel chan int, channelDigit chan int) {
	sum := 0
	for digit := range channelDigit {
		sum += digit * digit
	}
	channel <- sum
}

func calcCubesWithLoopChannel(channel chan int, channelDigit chan int) {
	sum := 0
	for digit := range channelDigit {
		sum += digit * digit * digit
	}
	channel <- sum
}

func channelTestWithLoop() {
	number := 589
	squareChannel := make(chan int)
	cubechChannel := make(chan int)
	channelDigitSquare := make(chan int)
	channelDigitCubes := make(chan int)
	go digits(number, channelDigitSquare, channelDigitCubes)
	go calcSquaresWithLoopChannel(squareChannel, channelDigitSquare)
	go calcCubesWithLoopChannel(cubechChannel, channelDigitCubes)
	squaresData := <-squareChannel
	cubechData := <-cubechChannel
	fmt.Println("Data Sum Digit With Channel Loop: ", squaresData+cubechData)
}
