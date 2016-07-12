// default.go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e9)
	boom := time.After(5e9)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e8)
		}
	}
}

/* Output:
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
BOOM!
*/
