package main

import "fmt"

func main() {

	mybill := newBill("mario's bill")

	fmt.Println(mybill) //returns {mario's bill map[] 0}

	fmt.Println(mybill.format()) //pie: ... $5.99 cake: ... $3.99 total:                    ...$9.98

}
