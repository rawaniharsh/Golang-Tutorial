package main

import (
	"fmt"
	"strings"
)

func main() {
	// create string using var
	var message1 = "Hello,"

	// create string using shorthand notation
	message2 := "Welcome to Programiz"

	fmt.Println(message1)
	fmt.Println(message2)

	// access first character
	fmt.Printf("%c\n", message1[0])

	//length of the string
	fmt.Println("Length of a string is:", len(message2))

	// concatenation using + operator
	result := message1 + " " + message2

	fmt.Println(result)

	//String Methods
	//the strings package provides various methods that can be used to perform different operations on strings.
	//Compare()	compares two strings
	// create three strings
	name1 := "Akash"
	name2 := "Rahul Singh"
	name3 := "Akash"

	// compare strings
	fmt.Println(strings.Compare(name1, name2)) // -1
	fmt.Println(strings.Compare(name2, name3)) // 1
	fmt.Println(strings.Compare(name1, name3)) // 0

	//Contains()	checks if a substring is present inside a string
	result2 := strings.Contains(name2, "Singh")
	fmt.Println(result2)

	//Replaces()	replaces a substring with another substring

	str := "I am Engineer"

	// replace am with was
	replacedStr := strings.Replace(str, "am", "was", 1)

	fmt.Println("New String:", replacedStr)

	/* ToUpper() - to change string to uppercase
	ToLower() - to change string to lowercase */

	//Split() - split strings method
	//The Split() method returns a slice of all the substrings.
	var message = "I Love Golang"

	// split string from space " "
	splittedString := strings.Split(message, " ")

	fmt.Println(splittedString)

	//Escape Sequence in Golang String
	escMsg := "My hobby is playing \"PUBG\" online"

	fmt.Println(escMsg)
}
