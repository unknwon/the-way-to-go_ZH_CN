// pointer_value.go
package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b1 B // b1 is value
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2 is pointer
	b2.change()
	fmt.Println(b2.write())
}

/* Output:
{1}
{1}
*/
