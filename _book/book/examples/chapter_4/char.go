// char.go
package main

import (
	"fmt"
)

func main() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2' 
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3)
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3)
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3)
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)
}
/* Ouput: 
65 - 946 - 1053236
A - β - 􁈴
41 - 3B2 - 101234
U+0041 - U+03B2 - U+101234
*/