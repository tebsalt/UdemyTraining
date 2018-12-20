package main

import "fmt"

func main() {
	i := 0 // initialization
	for {
		fmt.Println(i)
		if i >= 10 /*condition*/ {
			break //will kick us outside the loop
		}
		i++
	}
}
