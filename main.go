package main

import "fmt"

func main() {

	mybill := newBill("mario's bill")

	mybill.addItems("onion soup", 4.50)
	mybill.addItems("veg pie", 8.95)
	mybill.addItems("toffee pudding", 4.95)
	mybill.addItems("coffee", 3.25)

	mybill.updatedTip(10)

	//fmt.Println(mybill) //returns {mario's bill map[] 0}

	fmt.Println(mybill.format()) //pie: ... $5.99 cake: ... $3.99 total:                    ...$9.98

}
