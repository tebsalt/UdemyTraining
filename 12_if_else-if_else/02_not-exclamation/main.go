package main

import "fmt"

func main() {
	if !true {
		fmt.Println("This did not run")
	}
	if !false {
		fmt.Println("This ran")
	}
}

// the ! sign negates the statement after it
// so true becomes false and false becomes true
