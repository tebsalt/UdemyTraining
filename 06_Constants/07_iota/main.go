package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1*10)
	MB = 1 << (iota * 10) // 1 << (2*10)
	GB = 1 << (iota * 10) // 1 << (3*10)
	TB = 1 << (iota * 10) // 1 << (4*10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Println("%b\t", KB)
	fmt.Println("%d\n", KB)
	fmt.Println("%b\t", MB)
	fmt.Println("%d\n", MB)
	fmt.Println("%b\t", GB)
	fmt.Println("%d\n", GB)
	fmt.Println("%b\t", TB)
	fmt.Println("%d\n", TB)

}
