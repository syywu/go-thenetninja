package main

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
func (b bill) format() string {

}
