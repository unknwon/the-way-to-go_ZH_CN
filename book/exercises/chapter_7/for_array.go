package main

import "fmt"

func main() {
	var arr [15]int
	for i:=0; i < 15; i++ {
		arr[i] = i
	}
	fmt.Println(arr)   // [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
	

}