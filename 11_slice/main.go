package main

import "fmt"

/* Slice is a collection of similar types of data, just like arrays.

However, unlike arrays, slice doesn't have a fixed size. We can add or remove elements from the array. */

func main() {

	//Create slice in Golang
	numbers := []int{1, 2, 3, 4, 5}

	// print the slice
	fmt.Println("Numbers: ", numbers)

	//Adds Element to a Slice
	// add elements 5, 7 to the slice
	numbers = append(numbers, 5, 7)
	fmt.Println("Prime Numbers:", numbers)

	//add elements of one slice to another
	// create two slices
	evenNumbers := []int{2, 4}
	oddNumbers := []int{1, 3}

	// add elements of oddNumbers to evenNumbers
	evenNumbers = append(evenNumbers, oddNumbers...)
	fmt.Println("Numbers:", evenNumbers)

	// create two slices
	primeNumbers := []int{2, 3, 5, 7}

	// copy elements of primeNumbers to numbers
	copy(numbers, primeNumbers)

	// print numbers
	fmt.Println("Numbers:", numbers)

	//Find the Length of a Slice
	// find the length of the slice
	length := len(numbers)

	fmt.Println("Length:", length)

	// for loop that iterates through the slice
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	  }
}
