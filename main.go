package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// functions
func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

// n = name and is of type string

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v) //function takes in value
	}
}

// n is name of type slice with string inside
// function can be a param inside another function

func circleArea(r float64) float64 {
	return math.Pi * r * r //math.Pi is from package math. pi * radius squared
} // must state type before {} to specify which type to return

// return mulitple vales
func getInitials(n string) (string, string) { //returns 2 strings
	s := strings.ToUpper(n)           //get all names to be CAPS
	allNames := strings.Split(s, " ") //splits s into array when there is a space

	var initials []string //declare a varible of type slice called initals where appended initials to be stored
	for _, v := range allNames {
		initials = append(initials, v[:1]) //add initials that starts with 0 and ends at 1
	}

	if len(initials) > 1 {
		return initials[0], initials[0]
	}
	return initials[0], "_"
}

/* packages
var score = 99.5 //can only be found outside main(), inside the package main scope
// cannot use shorthand outside of functions
// scoreTwo := 50 */

// only one main func at the entry file. must called main to match with package
func main() {

	/* packages
	sayHello("mario")
	showScore()

	for _, v := range points {
		fmt.Println(v)
	} */

	// functions
	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	cycleNames([]string{"cloud", "barret", "tifa"}, sayGreeting)
	//Good morning cloud 		//f(v)
	// Good morning barret
	// Good morning tifa
	cycleNames([]string{"cloud", "barret", "tifa"}, sayBye) // invoke inside the cycleNames function
	// Goodbye cloud
	// Goodbye barret
	// Goodbye tifa

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 area is %0.3f & circle 2 area is %0.3f \n", a1, a2)

	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn1) //T L

	fn2, sn2 := getInitials("cloud strife")
	fmt.Println(fn2, sn2) //C S

	fn3, sn3 := getInitials("barret")
	fmt.Println(fn3, sn3) //B _

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

	// len(char) instead of char.length
	for i := 0; i < len(char); i++ {
		fmt.Println(char[i])
	}

	// for in example
	// uses := range instead of in
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

	nom := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, val := range nom {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue //continues the loop
			// continuing at pos 1
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break //breaks away from loop
			// breaking at pos 3 as it is > 2. nothing runs after this
		}
		fmt.Printf("the value at pos %v is %v \n", index, val)
		//the value at pos 0 is mario. this is the input as index is not equal to 0
		// the value at pos 2 is yoshi as index != 1

		// maps - like obj in JS. key : value pairs
		// specify key type in [] and value type before {}
		menu := map[string]float64{
			"soup":           4.99,
			"pie":            7.99,
			"salad":          6.99,
			"toffee pudding": 3.55,
		}

		fmt.Println(menu)        //map[pie:7.99 salad:6.99 soup:4.99 toffee pudding:3.55]
		fmt.Println(menu["pie"]) //returns the value = 7.99

		// looping maps
		for k, v := range menu {
			fmt.Println(k, "-", v) // soup - 4.99
			// k instead on index
		}

		// ints as key type
		phonebook := map[int]string{
			267584967: "mario",
			984759373: "luigi",
			845775485: "peach",
		}

		fmt.Println(phonebook)            //map[267584967:mario 845775485:peach 984759373:luigi]
		fmt.Println(phonebook[845775485]) //mario

	}
}
