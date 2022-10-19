package main

import "fmt"

/* A goroutine is a lightweight execution thread in the Go programming language and
a function that executes concurrently with the rest of the program.

Goroutines are incredibly cheap when compared to traditional threads as
the overhead of creating a goroutine is very low. Therefore,
they are widely used in Go for concurrent programming.

To invoke a function as a goroutine, use the go keyword.

Concurrent programs are able to run multiple processes at the same time.

With Goroutines, concurrency is achieved in Go programming. It helps two or more independent functions to run together.
Goroutines can be used to run background operations in a program.
It communicates through private channels so the communication between them is safer.
With goroutines, we can split one task into different segments to perform better. */

// create a function
func display(message string) {

	fmt.Println(message)
}

func main() {
 // call goroutine
 go display("Process 1")

 display("Process 2")
}
