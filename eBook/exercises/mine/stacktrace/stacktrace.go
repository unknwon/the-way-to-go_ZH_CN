package main

import (
	"runtime/debug"
)

// http://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces
func main() {

	test1()

}

func test1() {
	debug.PrintStack()
}
