/*
iteration_list.go:12: cannot define new methods on non-local type list.List
iteration_list.go:17: lst.Iter undefined (type *list.List has no field or method Iter)
---- Build file exited with code 1
*/
package main

import "container/list"

// cannot define new methods on non-local type list.List
// List iterator:
func (p *list.List) Iter() {
}

func main() {
	lst := new(list.List)
	for _ = range lst.Iter() {
	}
}




