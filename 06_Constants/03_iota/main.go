package main

import "fmt"

const (
	A = iota //0
	B = iota //1
	C = iota //2
)

func main() {
	fmt.Println(B)
	fmt.Println(A)
	fmt.Println(C)

}
