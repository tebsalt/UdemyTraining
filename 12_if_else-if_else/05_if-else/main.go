package main

import "fmt"

func main() {
	if false {
		fmt.Println("first print statement")
	} else {
		fmt.Println("second print statement")
	}
}

// if false "first print statement" wouldn't run but the else one
