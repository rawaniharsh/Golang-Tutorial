//lets write our first go program

// Print "Hello World!" message

package main

import "fmt"

func main() {
	//go print statement
	//fmt.Print()

	fmt.Print("Hello, ")
	fmt.Print("World!")

	//multiple values and variables at once by separating them with commas
	var name string = "Harsh"
	fmt.Print("Hello", name)

	//fmt.Println()
	//prints a new line at the end
	//a space is added between the values by default
	currentAge := 40

	fmt.Println("Hello")
	fmt.Println("World!")
	fmt.Println("Current Age:", currentAge)

	//fmt.Printf()
	//it formats the strings and print
	weight := 20.932
	fmt.Printf("Annual salary = %g\n", weight)

	//every data type has a unique format specifier
	//int - %d, float - %g, string - %s, bool - %t

	/* multiline commets, we can add here
	this is the syntax
	*/

}
