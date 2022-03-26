package main

import "fmt"

// only one main func at the entry file. must called main to match with package
func main() {
	var nameOne string = "mario"
	nameTwo := "luigi"
	// both are same. := shorthand cannot be used out

	fmt.Println(nameOne, nameTwo)

	// bits & memory

	var numOne int = 89
	var numTwo uint = 23
	// uint are only positive numbers
	var numThree int8 = -12
	// int8 ranges from -128 to 127
	// bits dictate the range of numbers that can be used.

	// floats

	var scoreOne float32 = 25.87
	var scoreTwo float64 = 763873548237.34
	// float64 provides higher precision
}
