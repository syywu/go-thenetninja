package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{}, //makes an empty map
		tip:   0,                    //sets tip to 0
	}
	return b
}

// b is a bill object of inital values

// format the bill- receiver function. limits the function only to bill object. (b bill) b of type bill
// pass pointers to reciver func so not making a copy everytime we call a func. if calling func a lot then making loads of copies but if passing in a pointer then only making a copy of pointer
// if bill obj is large and complex, making copies of that would be more work than just a pointer
func (b *bill) format() string {

	fs := "Bill Breakdown: \n" //fs = formated strings
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) //adding items to fs
		total += v                                     //adding the value to total
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total+b.tip) //-25 makes the val 25 characters long which makes the list more inline
	/* pie:                      ...$5.99
	   cake:                     ...$3.99
	   total:                    ...$9.98

	   +25 will make everything move to the right

	   adding : individually from %v, otherwise will turn out like:

	   pie                     : ...$5.99


	*/

	return fs

}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip //(*b).tip but go auto dereference it

}

// rule of thumb- whenever we are calling a method where we are updating a val, we pass in a pointer

// add items
func (b *bill) addItems(name string, price float64) {
	b.items[name] = price //pointer as we are chaning the items here
}
