package main

import "fmt"

//data types in go to is type of data associated with variables.
//i.e int, float, complex, string, bool, byte, rune
func main() {
	//integer data type
	var age int = 30
	age2 := 40

	fmt.Println(age, age2)

	//float data type
	var percentage1 float32 = 4504432432.45434232434

	// can store decimals with greater precision
	var percentage2 float64 = 409555343.0534553234323

	fmt.Println(percentage1, percentage2)

	//string data type
	//string is a sequence of characters
	var dogName string = "Tommy"
	fmt.Println(dogName)

	//boolean data type
	var isLogin bool = false
	fmt.Println(isLogin)
}
