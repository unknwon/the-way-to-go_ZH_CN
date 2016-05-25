package main

import (
	"fmt"
	"time"
)

func main() {
	suck(pump())
	time.Sleep(1e9)
//fmt.Println("total ", total)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

var total int
func suck(ch chan int) {
	go func() {
		for v := range ch {
//++ total
			fmt.Println(v)
//++total
		}
	}()
}
