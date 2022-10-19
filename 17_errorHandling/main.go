package main

import (
	"errors"
	"fmt"
)

/* When an error occurs, the execution of a program stops completely with a built-in error

we don't use try/catch to handle errors in Go. We can handle errors using:

New() Function
Errof() Function
*/

// function that checks day
func findDay(day string) error {

	// create a new error
	newError := errors.New("invalid day")

	// return the error if name is not Programiz
	if day != "Monday" {
		return newError
	}

	// return nil if there is no error
	return nil
}

func main() {
	message := "Hello"

	// create error using New() function
	myError := errors.New("WRONG MESSAGE")

	if message != "Programiz" {
		fmt.Println(myError)
	}

	// call the function
	err := findDay(message)

	// check if the err is nil or not
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid Day")
	}

}
