package main

import "fmt"

func main() {
	/* Switch statement in go tests the value of a variable and compares it
	with multiple cases. Once the case match is found, a block of statements
	associated with that particular case is executed.*/

	ageOfStudent := 12

	switch ageOfStudent {

	case 10:
		fmt.Println("Class 1")

	case 12:
		fmt.Println("Class 2")

	case 14:
		fmt.Println("Class 3")

	default:
		fmt.Println("Invalid age")
	}

	/* like C and Java, we don't need to use break after every case.
	This is because in Go, the switch statement terminates after the first matching case.*/

	//switch case with fallthrough
	switch ageOfStudent {

	case 10:
		fmt.Println("Class 1")

	case 12:
		fmt.Println("Class 2")
		fallthrough

	case 14:
		fmt.Println("Class 3")

	default:
		fmt.Println("Invalid age")
	}

	//Go switch with multiple cases
	dayOfWeek := "Sunday"

	switch dayOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}

}
