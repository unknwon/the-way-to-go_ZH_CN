package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4
type Stack [LIMIT]int

func main() {
	st1 := new(Stack)
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

// put value on first position which contains 0, starting from bottom
func (st *Stack) Push(n int) {
	for ix, v := range st {
		if v == 0 {
			st[ix] = n
			break
		}
	}
}

// take value from first position which contains !=0, starting from top
func (st *Stack) Pop() int {
	v := 0
	for ix:= len(st)-1; ix>=0; ix-- {
		if v=st[ix]; v!=0 {
			st[ix] = 0
			return v
		}
	}
	return 0
}

func (st Stack) String() string {
	str := ""
	for ix, v := range st {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(v) + "] "
	}
	return str
}