// factorial.go
package main

import (
	"fmt"
)

func main() {
	for i := uint64(0); i < uint64(30); i++ {
		fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
	}
}

/* unnamed return variables:
func Factorial(n uint64) uint64 {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}
*/

// named return variables:
func Factorial(n uint64) (fac uint64) {
	fac = 1
	if n > 0 {
		fac = n * Factorial(n-1)
		return
	}
	return
}

// int: correct results till 12!
// uint64: correct results till 21!
