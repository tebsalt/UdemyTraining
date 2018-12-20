package main

import "fmt"

var x int = 42

func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}

//anything declare in the package level scope (var x int = 42) is accessible inside scopes that are enclosed within the package level.
//
