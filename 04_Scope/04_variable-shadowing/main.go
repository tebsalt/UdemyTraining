package main

import "fmt"

func max(x int) int {
	//max isn't visible outside the package, cause its lower case and not upper cause
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) //max is now result, not the function

}

//don't do this; bad coding practice to shadow variables
//always have unique name for everything
