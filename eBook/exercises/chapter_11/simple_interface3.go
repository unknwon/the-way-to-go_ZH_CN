// simple_interface2.go
package main

import (
	"fmt"
)

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Set(u int) {
	p.i = u
}

type RSimple struct {
	i int
	j int
}

func (p *RSimple) Get() int {
	return p.j
}

func (p *RSimple) Set(u int) {
	p.j = u
}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(50)
		return it.Get()
	default:
		return 99
	}
	return 0
}

func gI(any interface{}) int {
	// return any.(Simpler).Get()   // unsafe, runtime panic possible
	if v, ok := any.(Simpler); ok {
		return v.Get()
	}
	return 0 // default value
}

/* Output:
6
60
*/

func main() {
	var s Simple = Simple{6}
	fmt.Println(gI(&s)) // &s is required because Get() is defined with a receiver type pointer
	var r RSimple = RSimple{60, 60}
	fmt.Println(gI(&r))
}

/* Output:
6
60
*/
