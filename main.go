package  main

import (
	"fmt"
)

func main()  {
	learnIf()
	learnSwitch()
	learnMltipleSwitch()
}

func learnIf()  {
	num := 10
	if num % 2 == 0 {
		fmt.Printf("Num %v là số chẵn", num)
	}else {
		fmt.Printf("Num %v là số lẽ", num)
	}
}

func  learnSwitch()  {
	num := 10
	switch num % 2 {
	case 0:
		fmt.Printf("Num %v là số chẵn\n",num)
		break
	case 1:
		fmt.Printf("Num %v là số lẻ\n",num)
		break
	default:
		break
	}
}

func learnMltipleSwitch()  {
	num := 10
	switch num % 2 {
	case 0,1:
		fmt.Printf("Num %v là số",num)
	default:
		break
	}
}