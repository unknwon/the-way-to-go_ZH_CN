package main

import "fmt"

type C struct {
	x float32
	int
	string
}

type A struct{ a int }
type B struct{ a, b int }

type C1 struct {
	A
	B
}

var c1 C1

func main() {
	//	c := C{3.14, 7, "hello"}
	//	fmt.Println(c.x, c.int, c.string) // output: 3.14 7 hello
	//	fmt.Println(c)                    // output: {3.14 7 hello}

	c1 = C1{A{1}, B{2, 3}}
	fmt.Println(c1)
	//	print(c1.a)
	println(c1.A.a)
	println(c1.B.a)
}
