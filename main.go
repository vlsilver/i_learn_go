package main

import "fmt"

func main() {
	//Go makes "copies" of values when passed into functions
	//Gruop A - value type: Strings, Ints, Floats, Booleans, Arrays, Structs
	//Group B - reference type(pointer): Slices, Maps, Functions
	name := "vlsilver"
	updateName(name)
	fmt.Println(name)
	menu := map[string]int{"pie": 1, "coffe": 2}
	updateMenu(menu)
	fmt.Println(menu)


}

func updateName(x string) {
	x = "vl"
}
func  updateMenu(m map[string]int)  {
	m["coffee"] = 4
}
