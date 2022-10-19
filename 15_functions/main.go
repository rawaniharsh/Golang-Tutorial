package main

import "fmt"

/* A function is a block of code that performs a task.
It can be called and reused multiple times.
You can pass information to a function and it can send information back. */

//we use the func keyword to create a function.
// define a function
func greet() {
	fmt.Println("Good Morning")
}

//Function with Parameters
func addNumbers(n1 int, n2 int) int {
	sum := n1 + n2
	// fmt.Println("Sum:", sum)

	//return values from function in golang
	return sum
}

//Return Multiple Values from Go Function
// function definition
func calculate(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	difference := n1 - n2

	// return two values
	return sum, difference
}

//An anonymous function is a function that was declared without any named identifier to refer to it.
//We assign the anonymous function to a variable and then use the variable name to call the function.
// anonymous function with arguments
var Add = func(n1, n2 int) {
	sum := n1 + n2
	fmt.Println("Sum is:", sum)
}

func main() {
	// function call
	greet()

	// pass parameters in function call
	result := addNumbers(21, 13)

	fmt.Println("Sum:", result)

	// function call
	sum, difference := calculate(21, 13)

	fmt.Println("Sum:", sum, "Difference:", difference)

	Add(7, 8)
}
