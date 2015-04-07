package main

import "fmt"

type obj interface{}

func main() {
	// define a generic lambda function mf:
	mf := func(i obj) obj {
		switch i.(type) {
			case int:
				return i.(int) * 2
			case string:
				return i.(string) + i.(string)
		}
		return i
	}
	
	res1 := mapFunc(mf, 0, 1, 2, 3, 4, 5)
	for _, v := range res1 {
		fmt.Println(v)
	}
	println()
	res2 := mapFunc(mf, "0", "1", "2", "3", "4", "5")
	for _, v := range res2 {
		fmt.Println(v)
	}
}

func mapFunc(mf func(obj) obj, list ...obj) ([]obj) {
	result := make([]obj, len(list))
	
	for ix, v := range list {
		result[ix] = mf(v)
	}
	
	// Equivalent:
	/*
	for ix := 0; ix<len(list); ix++ {
		result[ix] = mf(list[ix])
	}
	*/
	return result
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