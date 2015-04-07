package main

import "fmt"

func main() {
	var min, max int
	min, max = MinMax(78, 65)
	fmt.Printf("Minimum is: %d, Maximum is: %d\n", min, max)
}

func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {		// a = b or a < b
		min = b
		max = a
	}
	return
}