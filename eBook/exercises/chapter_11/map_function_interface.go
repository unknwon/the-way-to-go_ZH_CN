package main

import "fmt"

func main() {
	mf := func(a obj) obj {
		switch v := a.(type) {
		case int:
			return 2*v
		case string:
			return v+v
		}
		return a
	}

	isl := []obj{0, 1, 2, 3, 4, 5}
	res1 := Map(mf, isl)
	for _, v := range res1 {
		fmt.Println(v)
	}
	println()

	ssl := []obj{"0", "1", "2", "3", "4", "5"}
	res2 := Map(mf, ssl)
	for _, v := range res2 {
		fmt.Println(v)
	}
}

type obj interface {}

func Map(fn func(obj) obj, s []obj) (res []obj) {
	res = make([]obj, len(s))
	for i, value := range s {
		res[i] = fn(value)
	}
	return
}

/* Output:
0
2
4
6
8
10

00
11
22
33
44
55
*/
