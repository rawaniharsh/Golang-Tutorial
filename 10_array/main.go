package main

import "fmt"

/* Arrays are used to store multiple values of the same type in a single variable,
instead of declaring separate variables for each value.
var array_variable = [size]datatype{elements of array} */

func main() {
	// declare array variable of type integer
	// defined size [size=5]
	var arrayOfInteger = [5]int{1, 5, 8, 0, 3}

	fmt.Println(arrayOfInteger)

	/*Accessing array elements in Golang
	each element in an array is associated with a number. The number is known as an array index.
	*/

	books := [4]string{"Ramayana", "Mathematics", "Rd Sharma", "RS agarwal"}

	// access element at index 0
	fmt.Println(books[0])

	// access element at index 3
	fmt.Println(books[3])

	//Changing the array element in Go
	books[2] = "Science world"

	fmt.Println(books)
}
