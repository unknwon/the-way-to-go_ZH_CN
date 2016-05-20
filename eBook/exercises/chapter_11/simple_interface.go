// simple_interface.go
package main

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Put(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Put(u int) {
	p.i = u
}

func fI(it Simpler) int {
	it.Put(5)
	return it.Get()
}

//wrong
//func (it *Simpler) f2() int {
//	&it.Put(6)
//	return &it.Get()
//}

func main() {
	var s Simple
	fmt.Println(fI(&s))  // &s is required because Get() is defined with a receiver type pointer
//	var sr Simpler = s
//	fmt.Println(s.
}
// Output: 5