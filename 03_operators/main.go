//An operator is a symbol that tells the compiler to perform specific mathematical or logical manipulations

package main

import "fmt"

func main() {
	//Arithmetic operators
	//+(addition), -(subtraction), *(multiplication), /(division), %(modulo division)

	number1 := 10
	number2 := 20
	// + adds two variables
	sum := number1 + number2
	fmt.Printf("%d + %d = %d\n", number1, number2, sum)

	// - subtract two variables
	diff := number1 - number2
	fmt.Printf("%d - %d = %d\n", number1, number2, diff)

	// * multiply two variables
	prod := number1 * number2
	fmt.Printf("%d * %d is %d\n", number1, number2, prod)

	// / divide two integer variables
	quotient := number1 / number2
	fmt.Printf(" %d / %d = %d\n", number1, number2, quotient)

	// % modulo-divides two variables
	remainder := number1 % number2
	fmt.Println(remainder)

	//Increment and Decrement Operator in Go
	// increment of num by 1
	number1++
	fmt.Println(number1)

	// decrement of num by 1
	number1--
	fmt.Println(number1)

	//Compound Assignment Operators
	number1 += 6

	//Relational Operators
	// ==(equal to) !=(not equal to) > (greater than) < (lesser than) >= (greater than equal to) <= (less than equal to)

	// equal to operator
	result := (number1 == number2)

	fmt.Printf("%d == %d returns %t \n", number1, number2, result)

	// not equal to operator
	result = (number1 != number2)

	fmt.Printf("%d != %d returns %t \n", number1, number2, result)

	// greater than operator
	result = (number1 > number2)

	fmt.Printf("%d > %d returns %t \n", number1, number2, result)

	// less than operator
	result = (number1 < number2)

	fmt.Printf("%d < %d returns %t \n", number1, number2, result)

	//Logical Operators in Go
	//&& (Logical and) || (Logical or) !(logical not)

	number3 := 6
	
	// returns false because number1 > number2 is false
	result = (number1 > number2) && (number1 == number3)

	fmt.Printf("Result of AND operator is %t \n", result)
  
	// returns true because number1 == number3 is true
	result = (number1 > number2) || (number1 == number3)
  
	fmt.Printf("Result of OR operator is %t \n", result)
	
	// returns false because number1 == number3 is true
	result = !(number1 == number3);
  
	fmt.Printf("Result of NOT operator is %t \n", result)

}
