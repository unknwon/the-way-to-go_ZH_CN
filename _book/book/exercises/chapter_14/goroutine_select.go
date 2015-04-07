// Q20b_goroutine.go
package main

import (
	"fmt"
	"os"
)

func tel(ch chan int, quit chan bool) {
	for i:=0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func main() {
	var ok = true
	ch := make(chan int)
	quit := make(chan bool)

	go tel(ch, quit)
	for ok {
		select {
			case i:= <-ch:
				fmt.Printf("The counter is at %d\n", i)
			case <-quit:
				os.Exit(0)
		}
	}
}

