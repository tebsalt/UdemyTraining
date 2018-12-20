package main

import "fmt"

func main() {
	i := 0 // initialize
	for {
		i++ //called incremental, not post, cause its happening earlier
		//still doing the exact same thing, it could be anywhere
		if i%2 == 0 /*only odd numbers will be printed, cause it's final answer cant be 0*/ {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
