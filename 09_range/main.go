package main

import "fmt"

/* The range keyword is mainly used in for loops in order to
iterate over all the elements of a map, slice, channel, or an array */

func main() {
	// array of numbers
	numbers := [5]int{21, 24, 27, 30, 33}

	// use range to iterate over the elements of array
	for index, item := range numbers {
		fmt.Printf("numbers[%d] = %d \n", index, item)
	}

	//range keyword returns two items

	//range with string in Golang
	string := "Golang"
	fmt.Println("Index: Character")

	// i access index of each character
	// item access each character
	for i, item := range string {
		fmt.Printf("%d= %c \n", i, item)
	}

}
