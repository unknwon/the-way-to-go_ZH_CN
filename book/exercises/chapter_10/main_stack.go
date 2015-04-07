// Q15.go
package main

import (
	"fmt"
	"./stack/stack"
)

func main() {
	st1 := new(stack.Stack)
	fmt.Printf("%v\n", st1)
	st1.Push(3)
	fmt.Printf("%v\n", st1)
	st1.Push(7)
	fmt.Printf("%v\n", st1)
	st1.Push(10)
	fmt.Printf("%v\n", st1)
	st1.Push(99)
	fmt.Printf("%v\n", st1)
	p := st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
}
/* Output:
[0:3]
[0:3] [1:7]
[0:3] [1:7] [2:10]
[0:3] [1:7] [2:10] [3:99]
Popped 99
[0:3] [1:7] [2:10]
Popped 10
[0:3] [1:7]
Popped 7
[0:3]
Popped 3
*/
