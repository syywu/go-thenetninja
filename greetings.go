package main

import "fmt"

// functions at the top level of main file/ package can be used elsewhere, where it is package main

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("hello", n)
}

/* func showScore() {
	fmt.Println("You scored this many points:", score)
} */
