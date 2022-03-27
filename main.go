package main

import (
	"fmt"
	"sort"
	"strings"
)

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
	people[1] = "mario"
	// this will change bill to mario

	fmt.Println(days, len(days))
	fmt.Println(people, len(people))
	// len() to get the length of arr

	// slices (use arrays under the hood)

	var results = []int{28, 34, 56}
	// slices have no fixed elements
	results[2] = 43 //56 instead of 43

	results = append(results, 89)
	// can append with slices only!!! it adds 89 to the slices

	fmt.Println(results, len(results))

	// slice ranges
	rangeOne := people[1:4]  // doesn't include pos 4 element- inclusive of 1st but not 4th element
	rangeTwo := people[2:]   //includes the last element
	rangeThree := people[:3] //start from beginning until position 3

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

	// package- strings

	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))         //retruns true
	fmt.Println(strings.Contains(greeting, "bye"))           //returns false
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) //replace hello to hi. original str remains the same. str is not altered
	fmt.Println(strings.ToUpper(greeting))                   // HELLO THERE FRIENDS!
	fmt.Println(strings.Index(greeting, "ll"))               //looks for the index. returns 2 as that is the position of the first 1
	fmt.Println(strings.Split(greeting, ""))                 //splits the str and splits it with wahtever valued entered. returns an array

	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	// package sort

	minutes := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(minutes) //[20 25 30 35 45 50 60 75] sort alters the origina slice
	fmt.Println(minutes)

	index := sort.SearchInts(minutes, 30) //returns 2 as slice has been altered. new slice = sort version
	// if the number entered in searchints does not exists in slice, then will return an index out of the slice i.e. 8
	fmt.Println(index)

	characters := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(characters) //sorts characters to alphabetical order
	fmt.Println(characters)

	fmt.Println(sort.SearchStrings(characters, "bowser")) //returns index 0

	//loops

	// for keyword is used on go for loops, for all for loops, for in and while

	// while loop example
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++ // infinite loop without this
	}

	// for loop example
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	// for loop looping through slices
	char := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(char); i++ {
		fmt.Println(char[i])
	}

	// for in example
	for index, val := range char {
		fmt.Printf("the value at position %v is %v \n", index, val)
		val = "new string"
	}

	// if you dont want either index or val, replace it with _
	for _, val := range char {
		fmt.Print(val, ",") //returns mario,luigi,yoshi,peach,
		val = "new string"
	}

	// changing val in a loop does not mutate the original slice
	fmt.Println(char) //returns [mario luigi yoshi peach]

	// bool & conditions

	num := 45

	fmt.Println(num <= 50) //true
	fmt.Println(num >= 50) //false
	fmt.Println(num == 45) //true
	fmt.Println(num != 50) //true

	if num < 30 {
		fmt.Println("num is less than 30")
	} else if num < 40 {
		fmt.Println("num is less than 40")
	} else {
		fmt.Println("num is not less than 40")
	}
}
