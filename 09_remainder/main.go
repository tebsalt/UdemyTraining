package main

import "fmt"

func main() {
	x := 13 % 3 // the / sign divides it...
	fmt.Println(x)
	if x == 4 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

}

// 13 / 3 will be 4; but the % causes it to just print out
// the remainder
