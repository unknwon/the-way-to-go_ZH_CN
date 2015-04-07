// Q29_uniq.go
package main

import (
	"fmt"
)

var arr []byte = []byte{'a','b','a','a','a','c','d','e','f','g'}

func main() {
	arru := make([]byte,len(arr)) // this will contain the unique items
	ixu := 0 // index in arru
	tmp := byte(0)
	for _, val := range arr {
		if val!=tmp {
			arru[ixu] = val
			fmt.Printf("%c ", arru[ixu])	
			ixu++
		}
		tmp = val
	}
	// fmt.Println(arru)
}
