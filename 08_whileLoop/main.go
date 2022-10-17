package main

import "fmt"

func main() {
	/* A while loop helps when we don't know in advance how often code inside the loop should
	repeat. while loop repeats code until a condition is true. Go doesn't have a dedicated keyword for a while loop.
	However, we can use the for loop to perform the functionality of a while loop.*/

	number := 1

	for number <= 5 {
		fmt.Println(number)
		number++
	}
}
