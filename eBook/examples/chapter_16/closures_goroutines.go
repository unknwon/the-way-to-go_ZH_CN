// closures_goroutines.go
package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// version A:
	for ix := range values { // ix is index!
		func() {
			fmt.Print(ix, " ")
		}() // call closure, prints each index
	}
	fmt.Println()
	// version B: same as A, but call closure as a goroutine
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	time.Sleep(1e9)
	// version C: the right way
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	time.Sleep(1e9)
	// version D: print out the values:
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	time.Sleep(1e9)
}

/* Output:
0 1 2 3 4
4 4 4 4 4
1 0 3 4 2
0 1 2 4 3
*/
