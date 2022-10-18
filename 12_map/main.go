package main

import "fmt"

/* the map data structure stores elements in key/value pairs.
Here, keys are unique identifiers that are associated
with each value on a map.
*/

func main() {
	// create a map
	subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}

	fmt.Println(subjectMarks)

	// access value for key Java
	fmt.Println(subjectMarks["Java"])

	// change value of Golang to 90
	subjectMarks["Golang"] = 90

	// add element with key JavaScript
	subjectMarks["javaScript"] = 56

	// create a map
	squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}

	// for-range loop to iterate through each key-value of map
	for number, squared := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", number, squared)
	}

}
