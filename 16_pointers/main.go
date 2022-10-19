package main

import "fmt"

//Pointers reference a location in memory where a value is stored rather than the value itself.

// function definition with a pointer argument
func update(num *int) {

	// dereference the pointer
	*num = 30

}

// call by value
func callByValue(num int) {

	num = 30
	fmt.Println(num) // 30

}

// call by reference
func callByReference(num *int) {

	*num = 10
	fmt.Println(*num) // 10

}

func main() {
	var num int = 5
	// prints the value stored in variable
	fmt.Println("Variable Value:", num)

	// prints the address of the variable
	fmt.Println("Memory Address:", &num)

	// create the pointer variable
	var ptr *int = &num
	fmt.Println("Value of pointer is", ptr)

	var number = 55

	// function call
	update(&number)

	fmt.Println("The number is", number)

	/* When we change *num to 30 inside the update() function,
	the value of number inside the main() function is also changed.

	This is because num and number are referring to the same address in the memory.
	So, updating one updates the other as well.

	Call By Reference
	While passing pointers to a function, we are actually passing a reference (address) of the variable.
	Instead of working with the actual value, we are working with references like

	accessing value using reference
	changing value using reference
	*/

	// passing value
	callByValue(number)

	// passing a reference (address)
	callByReference(&number)
}
