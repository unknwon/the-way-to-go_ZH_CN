// function_filter2.go
package main

import "fmt"

type flt func(int) bool

// func isEven(n int) bool { if n%2 == 0 { return true }; return false }
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func filter(sl []int, f flt) (yes, no []int) {
	for _, val := range sl {
		if f(val) {
			yes = append(yes, val)
		} else {
			no = append(no, val)
		}
	}
	return
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	even, odd := filter(slice, isEven)
	fmt.Println("The even elements of slice are: ", even)
	fmt.Println("The odd elements of slice are: ", odd)
}

/*

*/
