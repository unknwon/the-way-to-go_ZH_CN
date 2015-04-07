// interface_nil.go
package main

import "fmt"

type Any interface {}
type Anything struct {}

func main() {
	any := getAny()
	if any == nil {
		fmt.Println("any is nil")
	} else {
		fmt.Println("any is not nil")
	}
/*
	// to get the inner value:
	anything := any.(*Anything) 
	if anything == nil {
		fmt.Println("anything is nil")
	} else {
		fmt.Println("anything is not nil")
	}
*/
}

func getAny() Any {
	return getAnything()
}

func getAnything() *Anything {
	return nil
}

/* Output:
any is not nil
WHY?
you would perhaps expect: any is nil,because getAnything() returns that
BUT:
the interface value any is storing a value, so it is not nil. 
It just so happens that the particular value it is storing is a nil pointer.
The any variable has a type, so it's not a nil interface, 
rather an interface variable with type Any and concrete value (*Anything)(nil).
To get the inner value of any, use: anything := any.(*Anything)
now anything contains nil !
*/
