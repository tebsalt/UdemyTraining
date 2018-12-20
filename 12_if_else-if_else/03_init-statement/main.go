package main

import "fmt"

func main() {
	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
}

// usually you're checking if "chocolate - b" is available, then "food" is printed
// it can be printed as if b {fmt.println(food)}
// food:="chocolate" is the initialization
// the 2nd b is an expression of the if statement and initialization is made to check what we're after
