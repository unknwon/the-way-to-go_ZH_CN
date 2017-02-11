// nexter.go
package main

import (
	"fmt"
)

type nexter interface {
	next() byte
}

func nextFew1(n nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		b[i] = n.next()
	}
	return b
}

func nextFew2(n *nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		b[i] = n.next() // compile error:  n.next undefined (type *nexter has no field or method next)
	}
	return b
}

func main() {
	fmt.Println("Hello World!")
}
