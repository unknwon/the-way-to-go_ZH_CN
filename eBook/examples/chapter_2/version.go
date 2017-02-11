package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("%s", runtime.Version())
}

// Output:
// go1.0.3 or go 1.1
