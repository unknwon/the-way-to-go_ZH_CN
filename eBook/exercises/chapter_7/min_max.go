// Q13_1_max.go
package main

import (
	"fmt"
	"math"
)

func main() {
	sl1 := []int{78, 34, 643, 12, 90, 492, 13, 2}
	max := maxSlice(sl1)
	fmt.Printf("The maximum is %d\n", max)
	min := minSlice(sl1)
	fmt.Printf("The minimum is %d\n", min)
}

func maxSlice(sl []int) (max int) {
	for _, v := range sl {
		if v > max {
			max = v
		}
	}
	return
}

func minSlice(sl []int) (min int) {
	// min = int(^uint(0) >> 1)
	min = math.MaxInt32
	for _, v := range sl {
		if v < min {
			min = v
		}
	}
	return
}

/* Output:
The maximum is 643
The minimum is 2
*/
