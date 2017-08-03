package main

import "fmt"

type C struct {
	x float32
	int
	string
}

func main() {
	c := C{3.14, 7, "hello"}
	fmt.Println(c.x, c.int, c.string) // output: 3.14 7 hello
	fmt.Println(c)                    // output: {3.14 7 hello}
}
