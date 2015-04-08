// struct_conversions.go
package main

import (
	"fmt"
)

type number struct {
	f	float32
}

type nr number   // alias type

func main() {
	a := number{5.0}
	b := nr{5.0}
	// var i float32 = b   // compile-error: cannot use b (type nr) as type float32 in assignment
	// var i = float32(b)  // compile-error: cannot convert b (type nr) to type float32
	// var c number = b    // compile-error: cannot use b (type nr) as type number in assignment
	// needs a conversion:
	var c = number(b)
	fmt.Println(a, b, c)
}
// output: {5} {5} {5}
