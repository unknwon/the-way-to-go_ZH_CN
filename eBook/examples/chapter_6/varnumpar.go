package main

import "fmt"

func main() {
	x := Min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7, 9, 3, 5, 1}
	x = Min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

func Min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

/*
The minimum is: 0
The minimum in the slice is: 1
*/
