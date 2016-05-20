// mult_inheritance.go
package main

import "fmt"

type Camera struct { } 

func (c *Camera) TakeAPicture() string { 
    return "Click"
}

func (c *Camera) off() {
	println("c off")
}

type Phone struct { } 

func (p *Phone ) Call() string { 
    return "Ring Ring"
}

// fuyc: let's bring chaos with methods of same signature
func (p *Phone) off() {
	println("p off")
}

// multiple inheritance
type CameraPhone struct {
    Camera 
    Phone  
}

func main() {
    cp := new(CameraPhone)  
    fmt.Println("Our new CameraPhone exhibits multiple behaviors ...")
    fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())     
	fmt.Println("It works like a Phone too: ", cp.Call()) 
//	cp.off() // ambiguous selector cp.off
}
/* Output:
Our new CameraPhone exhibits multiple behaviors ...
It exhibits behavior of a Camera:  Click
It works like a Phone too:  Ring Ring
*/