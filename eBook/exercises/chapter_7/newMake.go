package main

import "fmt"

func main() {
	var p *[]int = new([]int) // *p == nil; with len and cap 0
	//		fmt.Printf("len %d cap %d\n", len(p), cap(p))
	p = new([]int)
	fmt.Printf("%s", p)

	p2 := make([]int, 0)
	fmt.Printf("%s", p2)

	v := make([]int, 10, 50)
	fmt.Printf("%s", v)

	s := make([]byte, 5)
	fmt.Printf("len %d cap %d", len(s), cap(s))
	s = s[2:4]
	fmt.Printf("len %d cap %d", len(s), cap(s))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Printf(" %s %s ", s1, s2)
	s2[1] = 't'
	fmt.Printf(" %s %s ", s1, s2)
}
