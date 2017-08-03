// prod_cons.go
/* producer-consumer problem in Go */
package main

import "fmt"

var done = make(chan bool)
var msgs = make(chan int)

func produce() {
	for i := 0; i < 10; i++ {
		msgs <- i
	}
	done <- true
}

func consume() {
	for {
		msg := <-msgs
		fmt.Print(msg, " ")
	}
}

func main() {
	go produce()
	go consume()
	<-done
}

// Output: 0 1 2 3 4 5 6 7 8 9
