package main

import "fmt"

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7}
	mf := func(i int) int {
		return i * 10
	}
	/*
	result := mapFunc(mf, list)
	for _, v := range result {
		fmt.Println(v)
	}
	*/
	println()
	// shorter:
	fmt.Printf("%v", mapFunc(mf, list) )
}

func mapFunc(mf func(int) int, list []int) ([]int) {
	result := make([]int, len(list))
	for ix, v := range list {
		result[ix] = mf(v)
	}
	/*
	for ix := 0; ix<len(list); ix++ {
		result[ix] = mf(list[ix])
	}
	*/
	return result
}
