package main

import "fmt"

var x = 1

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

// increment adds 1 to x and returns x
// when increment is called, the 1st time, it'll be 1, and the 2nd time, it'll be 2
