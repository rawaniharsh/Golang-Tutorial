package main

import "fmt"

/*
With the defer function, you are able to call a function
but prevent execution until some other functions are executed.

For your Go application to panic, it has to come to a point where it cannot
run anymore because it doesn't know what to do. This can be triggered by
the compiler or manually in our code.
As for the recover keyword, it is closely related with panic.
When a program enters a panic state we have an option to help it
recover from that, we use the recover keyword.
*/

// recover function to handle panic
func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}

func division(num1, num2 int) {

	// execute the handlePanic even after panic occurs
	defer handlePanic()

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

func main() {
	// defer the execution of Println() function
	defer fmt.Println("Three")

	fmt.Println("One")
	fmt.Println("Two")

	/*When we use multiple defer in a program, the order of execution of
	the defer statements will be LIFO (Last In First Out). */

	/*We use the panic statement to immediately end the execution of the program.
	If our program reaches a point where it cannot be recovered due to some major errors,
	it's best to use panic. */

	//Recover in go programming
	/* The panic statement immediately terminates the program. However,
	sometimes it might be important for a program to complete its execution and get some required results.*/

	fmt.Println("Help! Something bad is happening.")
	panic("Ending the program")
	fmt.Println("Waiting to execute")

	division(4, 2)
	division(8, 0)
	division(2, 8)
}
