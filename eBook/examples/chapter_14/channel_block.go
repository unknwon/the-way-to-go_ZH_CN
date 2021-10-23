package main

import (
	"fmt"
	"time"
)

var cnt = 0

func main() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0

	time.Sleep(time.Second)
	fmt.Println(cnt) // prints 1

}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i // the channel will be block due to lack of consumer
		cnt++   // this code will only execute once
	}
}
