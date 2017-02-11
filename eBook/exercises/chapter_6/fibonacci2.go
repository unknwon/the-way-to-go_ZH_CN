package main

import "fmt"

func main() {
	pos := 4
	result, pos := fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, result)
	pos = 10
	result, pos = fibonacci(pos)
	fmt.Printf("the %d-th fibonacci number is: %d\n", pos, result)
}

func fibonacci(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		v1, _ := fibonacci(n - 1)
		v2, _ := fibonacci(n - 2)
		val = v1 + v2
	}
	pos = n
	return
}
