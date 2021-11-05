package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Là một slice chưa các byte(int8)
	//Để xác định chính xác từng phần tử trong chuỗi ta cần chuyển string về []rune
	//Chuỗi là bất biến nên chúng ta không thể thay đổi chuỗi
	//->Để thay đổi ta cần chuyển chuỗi về []rune, chúng ta thay đổi []slice
	//->Và chuyển nó lại về dạng string

	learnString := "vlsivỄr" // UTF-8

	for i := 0; i < len(learnString); i++ {
		fmt.Printf("%x -  %c\n", learnString[i], learnString[i])
	}

	// Đoạn code ở trên xảy ra lỗi
	//76 -  v
	//6c -  l
	//73 -  s
	//69 -  i
	//76 -  v
	//e1 -  á
	//bb -  »
	//84 - ... // something here
	//72 -  r
	// Mã Unicode UTF-8 của Ễ là e1 & bb & 84
	//-> khi in ra từng byte thì sẽ hiển thị không đúng.
	//-> Chúng ta dùng đến rune (int32)
	//->rune đại diện cho một điểm mã trong go, không quan trọng điểm mã chứa bao nhiêu byte.
	learnStringRune := []rune(learnString)
	fmt.Println(learnStringRune)
	for i := 0; i < len(learnStringRune); i++ {
		fmt.Printf("%x - %c\n", learnStringRune[i], learnStringRune[i])
	}
	//loop with string, cách dễ dàng để triển khai rune cho string
	for index, rune := range learnStringRune {
		fmt.Printf("%v - %c\n", index, rune)
	}
	//độ dài của String
	fmt.Println(utf8.RuneCountInString(learnString))

	// thay đổi giá trị của chuỗi
	// learnString[0] = 01 -- error
	learnStringRune[0] = []rune("a")[0]
	fmt.Println(string(learnStringRune))

}
