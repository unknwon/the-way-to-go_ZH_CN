// magic.go
package main

import "fmt"

type Base struct{}

func (Base) Magic() { fmt.Print("base magic ") }

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() { fmt.Println("voodoo magic") }

func main() {
	v := new(Voodoo)
	v.Magic()     
	v.MoreMagic() 
}
/* Output:
voodoo magic
base magic base magic 
*/