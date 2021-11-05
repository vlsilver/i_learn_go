package  main

import "fmt"

func main()  {
	learnIf()
}

func learnIf()  {
	num := 10
	if num % 2 == 0 {
		fmt.Printf("Num %v là số chẵn", num)
	}else {
		fmt.Printf("Num %v là số lẽ", num)
	}
}
