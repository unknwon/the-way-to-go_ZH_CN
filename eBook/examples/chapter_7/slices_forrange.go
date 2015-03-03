package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 4)

	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
}
/*
Slice at 0 is: 1
Slice at 1 is: 2
Slice at 2 is: 3
Slice at 3 is: 4
*/
