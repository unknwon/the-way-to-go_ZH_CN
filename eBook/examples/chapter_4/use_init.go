package main

import (
	"./trans"
	"fmt"
)

var twoPi = 2 * trans.Pi // decl computes twoPi

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}
