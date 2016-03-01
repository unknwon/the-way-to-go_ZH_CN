package main

import "fmt"

var fibs [50]int64

func main() {
	fibs[0] = 1
	fibs[1] = 1
	
	for i:= 2; i < 50; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	
	for i:=0; i < 50; i++ {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", i, fibs[i])
	}
}