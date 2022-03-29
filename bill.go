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
		items: map[string]float64{"pie": 5.99, "cake": 3.99}, //makes an empty map
		tip:   0,                                             //sets tip to 0
	}
	return b
}

// b is a bill object of inital values

// format the bill- receiver function. limits the function only to bill object. (b bill) b of type bill
func (b bill) format() string {

	fs := "Bill Breakdown: \n" //fs = formated strings
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) //adding items to fs
		total += v                                     //adding the value to total
	}

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total) //-25 makes the val 25 characters long which makes the list more inline
	/* pie:                      ...$5.99
	   cake:                     ...$3.99
	   total:                    ...$9.98

	   adding : individually from %v, otherwise will turn out like:

	   pie                     : ...$5.99


	*/

	return fs
}
