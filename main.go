package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// helper func so dont have to type out the same code everytime we want the user to type in something. just call this
// this func accpets user input which is of type string and type error
// accepts reader as arg as we dont want to type it out everytime. reader is of type *bufio.Reader
// accpets promp of type string as arg for different q's asked
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt) //print not println as we want everything to be on the same line
	// get the info
	input, err := r.ReadString('\n') //input as its more generic and grabs error too

	return strings.TrimSpace(input), err
}

//create user input instead of hardcoding codes
func creatBill() bill {
	// get user input info by using a reader to read info that a user types in
	reader := bufio.NewReader(os.Stdin)
	// the standard input is the terminal. new reader reads info from terminal

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n') //read str after the new line char is enetered
	// // \n is essentially the enter key. reader will read after user has typed in enter. whta they entered will be captured in the name var

	// // get rid of whitespace entered
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)
	// this is much easier in the future than type in all of the above. makes it much more reusable

	// storing whatever typed to b var
	b := newBill(name)
	fmt.Print("Created new bill - ", b.name, "\n")

	return b

}

// option for user to add items or tips, or even to save the bill
func promptOptions(b bill) {
	// create new local reader for this func
	reader := bufio.NewReader(os.Stdin)
	// press a to add items, press s to save, press t to add or update tips
	opt, _ := getInput("Choose option(a- add items, s- to save, t- to add tips): ", reader)
	fmt.Print(opt)

}

func main() {

	mybill := creatBill()
	promptOptions(mybill)

	fmt.Println(mybill) //{sam's  bill map[] 0}

	// mybill := newBill("mario's bill")

	// mybill.addItems("onion soup", 4.50)
	// mybill.addItems("veg pie", 8.95)
	// mybill.addItems("toffee pudding", 4.95)
	// mybill.addItems("coffee", 3.25)

	// mybill.updatedTip(10)
	/*
		Bill Breakdown:
		onion soup:               ...$4.5
		veg pie:                  ...$8.95
		toffee pudding:           ...$4.95
		coffee:                   ...$3.25
		tip:                      ...$10.00
		total:                    ...$31.65

	*/

	//fmt.Println(mybill) //returns {mario's bill map[] 0}

	// fmt.Println(mybill.format()) //pie: ... $5.99 cake: ... $3.99 total:                    ...$9.98

}
