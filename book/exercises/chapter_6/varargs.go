// Q10_varargs.go
package main

import (
	"fmt"
)

func main() {
	printInts()
	println()
	printInts(2, 3)
	println()
	printInts(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func printInts(list ...int) {
//	for i:=0; i<len(list); i++ {
//		fmt.Printf("%d\n", list[i])
//	}
	for _, v := range list {
		fmt.Printf("%d\n", v)
	}
}