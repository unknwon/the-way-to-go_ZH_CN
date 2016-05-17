package main

//import "fmt"

func main() {
	var arr = []int{1, 2, 3}
	res := sum(arr[:])
	print(res)
}

func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}
