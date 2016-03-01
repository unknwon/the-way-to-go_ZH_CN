// i_undefined.go
package main

import (
	"fmt"
)

func main() {
	var i int
	for i=0; i<10; i++ {
		fmt.Printf("%v\n", i)
	}
	fmt.Printf("%v\n", i)  //<-- compile error:  undefined i
}
