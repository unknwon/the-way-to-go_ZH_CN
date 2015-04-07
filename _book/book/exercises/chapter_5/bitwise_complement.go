package main

import "fmt"

func main() {
	for i:=0; i <= 10; i ++ {
			fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}
}
/* Output:
the complement of 0 is: -1
the complement of 1 is: -10
the complement of 10 is: -11
the complement of 11 is: -100
the complement of 100 is: -101
the complement of 101 is: -110
the complement of 110 is: -111
the complement of 111 is: -1000
the complement of 1000 is: -1001
the complement of 1001 is: -1010
the complement of 1010 is: -1011
*/