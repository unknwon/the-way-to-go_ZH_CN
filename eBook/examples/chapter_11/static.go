// static.go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

var r io.Reader

func main() {
	r = os.Stdin
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	f, _ := os.Open("test.txt")
	defer f.Close()
	r = bufio.NewReader(f)
	var s *bytes.Buffer = new(bytes.Buffer)
	r = s
	fmt.Println(s)
}
