package main

import (
	"fmt"
	"time"
)

var fibs [50]int64

func main() {
	fibs[0] = 1
	fibs[1] = 1

	begin := time.Now()
	for i := 2; i < 50; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	dur := time.Now().Sub(begin)
	fmt.Printf("time:%s\n", dur)

	for i := 0; i < 50; i++ {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", i, fibs[i])
	}
}
