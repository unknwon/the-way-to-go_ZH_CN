// size_int.go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 10
	size := unsafe.Sizeof(i)
	fmt.Println("The size of an int is: ", size)
}
// The size of an int is:  4
