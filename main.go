package main

import "fmt"

// only one main func at the entry file. must called main to match with package
func main() {

	// string
	var nameOne string = "mario"
	var nameTwo = "joe"
	nameThree := "luigi"
	// both are same. := shorthand cannot be used out
	var nameFour string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "bob"
	nameFour = "sam"

	fmt.Println(nameOne, nameFour)

	// bits & memory

	var numOne int = 89
	var numTwo uint = 23
	// uint are only positive numbers
	var numThree int8 = -12
	// int8 ranges from -128 to 127
	// bits dictate the range of numbers that can be used.

	println(numOne, numTwo, numThree)

	// floats

	var scoreOne float32 = 25.87
	var scoreTwo float64 = 763873548237.34
	// float64 provides higher precision
	scoreThree := 87483.4
	// := is automatically float64

	println(scoreOne, scoreTwo, scoreThree)

	// Print & format

	name := "Sam"
	age := 28

	fmt.Println("my name is", name, "and my age is", age)
	fmt.Printf("my name is %v and my age is %v \n", name, age)
	fmt.Printf("my name is %q and my age is %q \n", name, age)
	// order matters in printf. %_ = format specifiers, v stands for variable, q = quote
	fmt.Printf("age is a type of %T \n", age) //T = type
	fmt.Printf("you scored %f points! \n", 255.55)
	fmt.Printf("you scored %0.1f points! \n", 255.55)
	// f = floats. specify how many decimal point you want e.g. 0.1 = 1 decimal point, 0.2 = 2 etc

	// Sprintf (saved formatted strings)

	var str = fmt.Sprintf("my name is %v and my age is %v \n", name, age)
	println("the saved string is: ", str)

	// arrays & slices

	// var days [3]int = [3]int{20, 25, 35}
	var days = [3]int{20, 25, 35}
	// arrays have a fixed length so specify the numbers of elements in an array

	people := [4]string{"bob", "bill", "jill", "jon"}

	fmt.Println(days, len(days))
	fmt.Println(people, len(people))
	// len() to get the length of arr

}
