package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1) // tons of numbers appear
	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
