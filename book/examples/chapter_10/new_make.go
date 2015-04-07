// annoy1.go
package main

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// OK:
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1
	// not OK:
	z := make(Bar) // compile error: cannot make type Bar
	z.thingOne = "hello"
	z.thingTwo = 1
	// OK:
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"
	// not OK:
	u := new(Foo)
	(*u)["x"] = "goodbye" // !! panic !!: runtime error: assignment to entry in nil map
	(*u)["y"] = "world"
}
