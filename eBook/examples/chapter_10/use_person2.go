package main

import (
	"./person"
	"fmt"
)

func main() {
	p := new(person.Person)
	// error: p.firstName undefined (cannot refer to unexported field or method firstName)
	// p.firstName = "Eric"
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName()) // Output: Eric
}
