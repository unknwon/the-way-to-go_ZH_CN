package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// consumer:
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()
	// producer:
	for {
		select {
		case c <- 0: 
		case c <- 1:
		}
	}

}
