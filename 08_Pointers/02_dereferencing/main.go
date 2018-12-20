package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc0420080b8

	var b *int = &a
	fmt.Println(b)  // 0xc0420080b8
	fmt.Println(*b) // 43

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front b
	// this is known as dereferencing
	// the * is an opereator in this case
}
