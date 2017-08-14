// sort_persons.go
package main

import (
	"./sort"
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

type Persons []Person

func (p Persons) Len() int { return len(p) }

func (p Persons) Less(i, j int) bool {
	in := p[i].lastName + " " + p[i].firstName
	jn := p[j].lastName + " " + p[j].firstName
	return in < jn
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	p1 := Person{"Xavier", "Papadopoulos"}
	p2 := Person{"Chris", "Naegels"}
	p3 := Person{"John", "Doe"}
	arrP := Persons{p1, p2, p3}
	fmt.Printf("Before sorting: %v\n", arrP)
	sort.Sort(arrP)
	fmt.Printf("After sorting: %v\n", arrP)
}
