// blocking.go
// throw: all goroutines are asleep - deadlock!
package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	//out := make(chan int, 1) // solution 2
	// go f1(out)  // solution 1
	out <- 2
	go f1(out)
}
