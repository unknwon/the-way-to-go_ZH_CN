// static.go
package main

import (
	"io"
	"os"
	"bufio"
	"bytes"
	"fmt"
)

var r io.Reader

func main() {
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	f, _ := os.Open("test.txt")
	r = bufio.NewReader(f)
	var s *bytes.Buffer = new(bytes.Buffer)
	r = s
	fmt.Println(s)
}
