package main

import (
	"bufio"
	"fmt"
	"os"
)

//create user input instead of hardcoding codes
func creatBill() bill {
	// get user input info by using a reader to read info that a user types in
	reader := bufio.NewReader(os.Stdin)
	// the standard input is the terminal. new reader reads info from terminal

}

func main() {

	mybill := creatBill()

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

	fmt.Println(mybill.format()) //pie: ... $5.99 cake: ... $3.99 total:                    ...$9.98

}
