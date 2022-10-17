package main

import (
	"fmt"
)

func main() {
	//Type casting is the process of converting the value of one data type (integer, float, string) to another data type.
	// variable of float type
	var floatValue float32 = 9.8

	// convert float to int
	var intValue int = int(floatValue)

	//manually performing the type conversion, this is called explicit type casting in Go.
	fmt.Println(floatValue, intValue)

	//go doesn't support implicit type casting.
}
