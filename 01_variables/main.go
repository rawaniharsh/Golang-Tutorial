//a variable is a container that is used to store data
//Go is a strongly, statically typed language

package main

import "fmt"

func main() {

	//Assign Values to Go Variables
	//Method 1
	var name string = "Harsh"

	//Method 2
	var name1 = "Harsh"

	//Method 3
	name2 := "Harsh"

	//unsed variables gives error in go
	fmt.Println(name, name1, name2)

	//Changing Value of a Variable
	// initial value
	age := 10
	fmt.Println("Initial value", age)

	// change variable value
	age = 100
	fmt.Println("Changed value", age)

	//we cannot change the type of variables after it is declared.
	// cost := 10

	// Error code
	// assign string data
	// cost = "Hello"

	//Define multiple variables at Once
	var class, rollNo = "tenth", 22
	fmt.Println(class, rollNo)

	//constants in go
	// const eartArea = 8765545456576 // initial value

	// Error! Constants cannot be changed
	// eartArea = 2997924607867698987

}
