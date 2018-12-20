package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}

//to scan a name for login
//no ln in the first print, so the name is entered in the first line and not a new one beneath
