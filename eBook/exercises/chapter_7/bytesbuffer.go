package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("hi ")
	buf.WriteString(" you ")
	fmt.Print(buf.String())
}
