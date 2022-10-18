package main

import "fmt"

/* A struct is used to store variables of different data types.
Suppose we want to store the name and age of a person. We can create two variables: name and age and store value.

However, suppose we want to store the same information of multiple people.

In this case, creating variables for a person might be a tedious task. We can create a struct that stores the name and age to overcome this.

And, we can use this same struct for every person.
*/

func main() {
	//create a struct
	type Student struct {
		name   string
		rollNo int
	}

	var akash Student

	akash = Student{
		name: "Akash",
		rollNo: 2,
	}

	//create instance and define the value of name and age
	rahul := Student{"Rahul", 1}

	fmt.Println(rahul)

	// access the rollNo of the struct
	fmt.Println("Roll no:", rahul.rollNo)

	fmt.Println("Roll no:", akash.rollNo)
}
