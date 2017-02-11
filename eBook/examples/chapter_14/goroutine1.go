package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	// longWait()
	go longWait()
	// shortWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9) // sleep works with a Duration in nanoseconds (ns) !
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}
