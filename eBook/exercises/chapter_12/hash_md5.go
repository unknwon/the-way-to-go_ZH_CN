// hash_md5.go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	hasher := md5.New()
	b := []byte{}
	io.WriteString(hasher, "test")
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
}

/* Output:
Result: 098f6bcd4621d373cade4e832627b4f6
Result: [9 143 107 205 70 33 211 115 202 222 78 131 38 39 180 246]
*/
