package retangle

import ("math")

var retangle = 0

// init sẽ được chạy khi package được import
func init() {
	println("Hàm init() của package rectangle được gọi")
}

// Area Đây là một exported name
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal Đây là một exported name
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len*len) + (wid*wid))
	return diagonal
}
